package git4humans

import (
	"fmt"
	"os"
)

var Command = "gh"

func Handle() {
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
	case "push":
		Push()
	case "sync":
		Sync()
	case "branch":
		Branch()
	case "switch":
		SwitchBranch()
	case "repo":
		Repo()
	case "pr":
		Pr()
	case "user":
		User()
	case "help":
		Help()
	default:
		if Contains(GitCommands, command) {
			Git(command, args...)
		} else {
			fmt.Printf(`Error: '%[1]s' is not a valid command. See %[2]s help.
            `, command, Command)
		}
	}
}
