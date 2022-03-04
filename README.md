# Diary

A simple app written in Golang to log your day-to-day activities

![build status](https://github.com/b-turchyn/diary/actions/workflows/go.yml/badge.svg)

## Configuration

Copy `.diary.yaml.example` to your home directory to configure the location of the database.
**Without this, your diary will always write to the current directory.** This means that your
entries might end up in different places and you won't realize it.

**Known issue:** directory expansion (such as using `~` for the home directory or enviroment
variables) does not work and will be treated as literal characters.

## Usage

After [building](#building), run either `diary.exe` on Windows or `diary` on MacOS or Linux.
Running without any commands will provide the available options.

For all commands that you input data in (`learn`, `log`, and `mistake`), input can either be
passed in right on the command line, or the command can be run interactively. The two commands
are equivalent.

Non-interactive:
```bash
./diary learn The cake is a lie
Recorded: The cake is a lie
```

Interactive:
```
./diary learn
What did you learn? The cake is a lie
Recorded: The cake is a lie
```

### Tracking What You've Learned

The `learn` command is used for tracking interesting facts you learned today.

### Tracking What You're Doing

The `log` command is used for tracking what you currently are doing.

### Tracking Your Mistakes

The `mistake` command is used for tracking mistakes you've made so you can learn from them.

If you've got a bit more of a pottymouth, there's an alias for you to track your `f***up`s.

### What Did You Do Today?

The `today` command gives you a Markdown-formatted list of everything entered for today.

### Time Travelling

For all commands, the following options can be passed on the command line to add entries for a
previous day:
```
  -d, --date string       Use a specific date (--time is required with this)
  -t, --time string       Use a specific time
```

For example `./diary log -t 13:50` logs an entry at 1:50pm local time, and
`./diary today -d 2022-02-28` gets you your log entries for February 28th, 2022.

## Shell Aliases

To make interacting with `diary` easier for you, consider the following aliases in your `~/.bashrc`
or `~/.zshrc` (`diary` is located at `~/scripts/diary`, but this can be changed with the `q` alias:

```bash
alias yesterday='date -v-1d +%F'

alias q='~/scripts/diary'
alias qm='q mistake'
alias qe='q learn'
alias ql='q log'
alias qly='ql --date=$(yesterday) -t'
alias qt='q today'
alias qty='qt --date=$(yesterday)'
```

## Building

```bash
go get github.com/b-turchyn/diary
cd $(go env GOPATH)/src/github.com/b-turchyn/diary
go build
```

## FAQ

* Q: Why build this?<br>
  A: Because I wanted to experiment with Golang more.
* Q: Aren't there better ways to track this?<br>
  A: Almost certainly.
* Q: Isn't a text file just as good?<br>
  A: Probably.
* Q: Should I use this program?<br>
  A: Up to you. If you live on the command line, then a few aliases set up makes this pretty
     convenient. Otherwise, it's probably not for you.
