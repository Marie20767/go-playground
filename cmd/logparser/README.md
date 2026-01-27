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
- make it so the path to the logs is configurable via a command line flag -path=/logs/mylog.log
- query by timestamp and message (partial match)