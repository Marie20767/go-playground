package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

var ErrInvalidQuery = errors.New("invalid query param")

func main() {
	if err := run(); err != nil {
		log.Println("server closed", err)
		os.Exit(1)
	}
}

func run() error {
	if err := parseLogs("service=auth level=warning"); err != nil {
		return err
	}

	return nil
}

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Service   string    `json:"service"`
	Message   string    `json:"message"`
}

func parseLogs(query string) error {
	logs := []*Log{}
	// TODO: Update parseLogs() to loop through all /logs files instead of 1 hardcoded path
	filePath := "logs/mylog.log"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine, err := parseLogLine(line)
		if err != nil {
			return err
		}
		log, err := filterLog(query, parsedLine)
		if err != nil {
			return err
		}
		if log != nil {
			logs = append(logs, parsedLine)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	sortLogs(logs)
	for _, log := range logs {
		bytes, err := json.Marshal(log)
		if err != nil {
			return err
		}

		fmt.Println(string(bytes))
	}

	return nil
}

func sortLogs(logs []*Log) {
	slices.SortFunc(logs, func(a, b *Log) int {
		if a.Timestamp.Before(b.Timestamp) {
			return -1
		} else if a.Timestamp.After(b.Timestamp) {
			return 1
		}
		return 0
	})
}

func parseLogLine(line string) (*Log, error) {
	var parsed Log
	err := json.Unmarshal([]byte(line), &parsed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

var getters = map[string]func(*Log) string{
	"service": func(l *Log) string { return l.Service },
	"level":   func(l *Log) string { return l.Level },
}

func filterLog(query string, log *Log) (*Log, error) {
	params := strings.SplitSeq(query, " ")
	for param := range params {
		kv := strings.Split(param, "=")
		if len(kv) < 2 {
			return nil, ErrInvalidQuery
		}
		key, value := kv[0], kv[1]
		get, ok := getters[key]
		if !ok {
			return nil, ErrInvalidQuery
		}

		if value != get(log) {
			return nil, nil
		}
	}

	return log, nil
}
