package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type LogEntry struct {
	RawLine      string    `json:"-"`
	Timestamp    time.Time `json:"-"`
	RawTimestamp string    `json:"timestamp"`
	Level        string    `json:"level"`
	Service      string    `json:"service"`
	Message      string    `json:"message"`
}

type Query struct {
	Key   string
	Value string
}

func main() {
	path := flag.String("path", "logs", "path to the directory containing the log files")
	query := flag.String("query", "", "query to filter the logs by e.g. service=auth")
	flag.Parse()

	if err := run(*path, *query); err != nil {
		log.Println("log parsing failed", err)
		os.Exit(1)
	}
}

func run(path, query string) error {
	queries, err := parseQueries(query)
	if err != nil {
		return fmt.Errorf("invalid query: %v", err)
	}
	entries, err := parseLogFiles(path, queries)
	if err != nil {
		return fmt.Errorf("failed to parse log files: %v", err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp.Before(entries[j].Timestamp)
	})

	for _, entry := range entries {
		if matchesQueries(queries, entry) {
			fmt.Println(entry.RawLine)
		}
	}

	return nil
}

func parseQueries(query string) ([]Query, error) {
	parsedQueries := []Query{}
	if query == "" {
		return parsedQueries, nil
	}

	queries := strings.FieldsSeq(query)
	for query := range queries {
		parts := strings.SplitN(query, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid query format: %s (expected key=value)", parts)
		}

		parsedQueries = append(parsedQueries, Query{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	return parsedQueries, nil
}

func parseLogFiles(path string, queries []Query) (entries []LogEntry, err error) {
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || filepath.Ext(path) != ".log" {
			return nil
		}

		fileEntries, err := parseFile(path)
		if err != nil {
			return err
		}
		entries = append(entries, fileEntries...)

		return nil
	})

	return entries, err
}

func parseFile(filePath string) (entries []LogEntry, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read log file %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry LogEntry

		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &entry)
		if err != nil {
			return nil, fmt.Errorf("failed to parse log line %v:", err)
		}

		timestamp, err := time.Parse(time.RFC3339, entry.RawTimestamp)
		if err != nil {
			return nil, fmt.Errorf("failed to parse timestamp %v:", err)
		}
		entry.RawLine = line
		entry.Timestamp = timestamp

		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse file %v:", err)
	}

	return entries, nil
}

func matchesQueries(queries []Query, entry LogEntry) bool {
	for _, query := range queries {
		fn, ok := paramFilters[query.Key]
		if !ok {
			return false
		}

		match, err := fn(entry, query.Value)
		if !match || err != nil {
			return false
		}
	}

	return true
}

var paramFilters = map[string]func(LogEntry, string) (bool, error){
	"service": func(l LogEntry, q string) (bool, error) {
		return strings.EqualFold(l.Service, q), nil
	},
	"level": func(l LogEntry, q string) (bool, error) {
		return strings.EqualFold(l.Level, q), nil
	},
	"message": func(l LogEntry, q string) (bool, error) {
		return strings.Contains(strings.ToLower(l.Message), strings.ToLower(q)), nil
	},
	"after": func(l LogEntry, q string) (bool, error) {
		t, err := time.Parse(time.RFC3339, q)
		if err != nil {
			return false, err
		}
		return l.Timestamp.After(t), nil
	},
	"before": func(l LogEntry, q string) (bool, error) {
		t, err := time.Parse(time.RFC3339, q)
		if err != nil {
			return false, err
		}
		return l.Timestamp.Before(t), nil
	},
}
