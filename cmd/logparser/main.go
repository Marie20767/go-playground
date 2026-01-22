package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

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
		logs = append(logs, parsedLine)
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

func filterLog(query string, log *Log) (*Log, error) {
	// params := string[2]{}
	// loop through query
	// 		validate every key & value per query param (key=value)
	//		append to params
	//		[["service", "auth"], ["level", "error"]]

	// loop through params
	//		for every param, check if it's available in the log
	//			if yes: check value matches
	//						if no match: return nil
	//						if match: continue loop

	return log, nil

	// query example: "service=auth level=warning"
	// line example: {"timestamp": "2025-10-28T10:00:01Z", "service": "payments", "level": "info", "message": "Payment completed"}
}
