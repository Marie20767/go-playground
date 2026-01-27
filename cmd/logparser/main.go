package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
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
	query := flag.String("query", "", "")
	flag.Parse()
	if err := parseLogs(*query); err != nil {
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

// TODO: query by timestamp (after={timestamp} before={timestamp})
// TODO: query by partial message match

func parseLogs(query string) error {
	queryParams, err := parseQueryParams(query)
	if err != nil {
		return err
	}

	for i, param := range queryParams {
		key, value := param[0], param[1]
		if key == "path" {
			queryParams = append(queryParams[:i], queryParams[i+1:]...)
			fileLogs, err := parseFile(value, queryParams)
			if err != nil {
				return err
			}
			printSortedLogs(fileLogs)
			return nil
		}
	}

	err = filepath.WalkDir("logs", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			fileLogs, err := parseFile(path, queryParams)
			if err != nil {
				return err
			}
			if fileLogs != nil {
				printSortedLogs(fileLogs)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func printSortedLogs(logs []*Log) error {
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

func parseFile(filePath string, queryParams [][]string) ([]*Log, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	filteredLogs := []*Log{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine, err := parseLogLine(line)
		if err != nil {
			return nil, err
		}

		log, err := filterLog(queryParams, parsedLine)
		if err != nil {
			return nil, err
		}
		if log != nil {
			filteredLogs = append(filteredLogs, log)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return filteredLogs, nil
}

func parseQueryParams(query string) ([][]string, error) {
	params := strings.SplitSeq(query, " ")
	parsedParams := [][]string{}

	for param := range params {
		kv := strings.Split(param, "=")
		if len(kv) < 2 {
			return nil, ErrInvalidQuery
		}

		parsedParams = append(parsedParams, kv)
	}

	return parsedParams, nil
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
	var parsed *Log
	err := json.Unmarshal([]byte(line), &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

var paramFilters = map[string]func(*Log, string) (bool, error){
	"service": func(l *Log, q string) (bool, error) {
		return strings.EqualFold(l.Service, q), nil
	},
	"level": func(l *Log, q string) (bool, error) {
		return strings.EqualFold(l.Level, q), nil
	},
	"message": func(l *Log, q string) (bool, error) {
		return strings.Contains(strings.ToLower(l.Message), strings.ToLower(q)), nil
	},
	"after": func(l *Log, q string) (bool, error) {
		t, err := time.Parse(time.RFC3339, q)
		if err != nil {
			return false, err
		}
		return l.Timestamp.After(t), nil
	},
	"before": func(l *Log, q string) (bool, error) {
		t, err := time.Parse(time.RFC3339, q)
		if err != nil {
			return false, err
		}
		return l.Timestamp.Before(t), nil
	},
}

func filterLog(params [][]string, log *Log) (*Log, error) {
	for _, p := range params {
		qKey, qVal := p[0], p[1]
		fn, ok := paramFilters[qKey]
		if !ok {
			return nil, ErrInvalidQuery
		}

		match, err := fn(log, qVal)
		if err != nil || !match {
			return nil, err
		}
	}
	return log, nil
}
