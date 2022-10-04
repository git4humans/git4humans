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
	case "new", "n":
		New()
	case "start", "s":
		Start()
	case "refresh":
		Refresh()
	case "+":
		Add()
	case "-", "del", "delete":
		Delete()
	case "rename", "ren":
		Move()
	case "move", "mv":
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
	case "status", "st":
		Status()
	case "log", "lg":
		Log()
	default:
		if contains(command, GitCommands) {
			Git(command, args...)
		} else {
			fmt.Println()
			fmt.Printf(`Error: '%[1]s' is not a valid command. See %[2]s help.
            `, command, Command)
		}
	}
}

func IsHelp() bool {
	args := os.Args[2:]
	isHelp := contains("--help", args) || contains("-h", args)

	return isHelp
}

func HasOption(option string, args []string) bool {
	index := indexOf(option, args)
	hasOption := index >= 0

	return hasOption
}

func RemoveOption(option string, args []string) []string {
	return remove(option, args)
}

func HasFlag(key string, args []string) bool {
	flag := GetFlag(key, args)
	hasFlag := len(flag) > 0

	return hasFlag
}

func GetFlag(key string, args []string) string {
	index := indexOf(key, args)
	indexFlag := index + 1

	hasFlag := index >= 0 && indexFlag < len(args)

	if hasFlag {
		return args[indexFlag]
	} else {
		return ""
	}
}

func RemoveFlag(key string, args []string) []string {
	index := indexOf(key, args)
	indexFlag := index + 1

	if index >= 0 {
		if indexFlag < len(args) {
			return append(args[:index], args[indexFlag+1:]...)
		} else {
			return append(args[:index], args[index+1:]...)
		}
	} else {
		return args
	}
}
