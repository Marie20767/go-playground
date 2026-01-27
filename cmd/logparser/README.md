# Log Parser

## Requirements

input log files are structued like:
```
{"timestamp": "2025-10-28T10:00:00Z", "service": "auth", "level": "error", "message": "Invalid token"}
{"timestamp": "2025-10-28T10:00:01Z", "service": "payments", "level": "info", "message": "Payment completed"}
{"timestamp": "2025-10-28T10:00:02Z", "service": "auth", "level": "warning", "message": "Slow login"}
```

Setup:
- Create a few .log files in directory called /logs, they can also be nested, e.g:
  /logs/mylog.log
  /logs/nested/somenestedlog.log
- add some log lines into them like the example above

The command-line tool should:
1. Load all logs from the /logs directory
2. Support queries like:
   - `service=auth`
   - `level=error`
   - `service=auth level=warning`
   (e.g. `service=auth` query would filter out anything that does not have "service": "auth")
3. Print results in chronological order

e.g. should be able to run something like:
```
go run main.go -query="service=auth level=error"
```
and then the log lines that match that query should be logged

Note - you can use the `bufio.NewScanner` in golang to efficiently read a file line by line


### Extra steps for after version 1 is working:
- make it so the path to the logs is configurable via a command line flag
- query by timestamp and message (partial match)

## Detailed plan:
1. Create cmd folder and /cmd/logparser folder
2. Create 2 new log files under /logs (nested/unnested)
3. Create main.go file with basic error handling and call parseLogs()
4. In main.go, start with basic parseLogs() func that:
        1. reads 1 log file for now (hard code the path to that file for now)
        2. prints every raw log line of that file
        4. prints in chronological order
5. Update parseLogs() to accept query param:
    - for now, pass query string to parseLogs (e.g. parseLogs("service=auth level=warning"))
    - break query string up into service & level
    - filter by service & level (need validation here!)
    - check if it prints correctly by running `go run main.go`
6. Update parseLogs() to loop through all /logs files instead of 1 hardcoded path
7. Accept CLI arguments for parseLogs()
8. Make it so the path to the logs is configurable by accepting something like `go run main.go -query="service=auth level=error path=customlogs"`
9. Update parseLogs() to query by timestamp and message (e.g. after={timestamp})