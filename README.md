# Go notifier

Tool to send notification after some time.
Depends on libnotify.

## Usage:

```
git clone https://github.com/soqet/gon.git
go build -o gon .
./gon <time> [summary] [text]
```
time in [go's duration format](https://pkg.go.dev/time#ParseDuration)
