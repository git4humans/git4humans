package git4humans

import (
	"fmt"
	"os"
)

var Command = "gt"

func Handle() {
	Command = os.Args[0]
	args := os.Args[1:]

	if len(args) > 0 {
		Execute()
	} else {
		Help()
	}
}

func Execute() {
	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "new":
		New()
	case "start":
		Start()
	case "+":
		Add()
	case "-":
		Delete()
	case "del":
		Delete()
	case "delete":
		Delete()
	case "rename":
		Move()
	case "move":
		Move()
	case "copy":
		Copy()
	case "save":
		Save()
	case "publish":
		Publish()
	case "pub":
		Publish()
	case "sync":
		Sync()
	case "branch":
		Branch()
	case "switch":
		SwitchBranch()
	case "repo":
		Repo()
	case "user":
		User()
	case "help":
		Help()
	case "pr":
		Pr()
	case "status":
		Status()
	case "s":
		Status()
	case "l":
		Log()
	default:
		if contains(GitCommands, command) {
			Git(command, args...)
		} else {
			fmt.Printf(`Err: '%[1]s' is not a valid command. See %[2]s help.
            `, command, Command)
		}
	}
}

func IsHelp() bool {
	args := os.Args[2:]
	isHelp := contains(args, "--help") || contains(args, "-h")

	return isHelp
}
