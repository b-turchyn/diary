# Diary

A simple app written in Golang to log your day-to-day activities

## Usage

After [building](#building), run either `diary.exe` on Windows or `diary` on MacOS or Linux.
Running without any commands will provide the available options.

### Configuration

Copy `.diary.yaml.example` to your home directory to configure the location of the database.
**Without this, your diary will always write to the current directory.** This means that your
entries might end up in different places and you won't realize it.

**Known issue:** directory expansion (such as using `~` for the home directory or enviroment
variables) does not work and will be treated as literal characters.

### Adding Old Entries

For all commands, the following options can be passed on the command line to add entries for a
previous day:
```
  -d, --date string       Use a specific date (--time is required with this)
  -t, --time string       Use a specific time
```

Example: `./diary log -t 13:50` Logs an entry at 1:50pm local time.


## Building

```bash
go get github.com/b-turchyn/diary
cd $(go env GOPATH)/src/github.com/b-turchyn/diary
go build
```
