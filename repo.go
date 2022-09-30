package git4humans

import (
	"fmt"
	"os"
	"strings"
)

func Repo() {
	if IsHelp() {
		RepoUsage()
		return
	}

	args := os.Args[2:]

	if len(args) > 0 {
		command := args[0]

		if command == "rename" {
			rename()
		} else if command == "show" {
			show()
		} else if command == "prune" {
			prune()
		} else if command == "url" {
			url()
		} else if command == "delete" {
			delete()
		} else if command == "del" {
			delete()
		} else if command == "-" {
			delete()
		} else {
			add()
		}
	} else {
		list()
	}
}

func RepoUsage() {
	fmt.Println()
}

func HasRepo(name string) bool {
	repos := GitStr("remote", "-v")
	return strings.Contains(repos, name)
}

func NoRepo(name string) bool {
	return !HasRepo(name)
}

func add() {
	command := os.Args[2]
	name := "origin"
	url := ""

	var args []string

	if command == "+" {
		args = os.Args[3:]
	} else {
		args = os.Args[2:]
	}

	if len(args) >= 2 {
		name = args[0]
		url = args[1]
	} else if len(args) >= 1 {
		url = args[0]
	}

	if len(url) > 0 {
		if HasRepo(name) {
			GitStr("remote", "rm", name)
		}

		Git("remote", "add", name, url)
	}

	list()
}

func delete() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	if HasRepo(name) {
		message := GitStr("remote", "rm", name)
		error := strings.Contains(message, "fatal:")

		if error {
			fmt.Println(message)
		} else {
			fmt.Println("Remote repository " + name + " has been removed.")

			Git("remote", "-v")
		}
	} else {
		fmt.Println("Err: remote repository " + name + " is not found.")
		fmt.Println()

		Git("remote", "-v")
	}
}

func rename() {
	args := os.Args[3:]

	oldName := "origin"
	newName := ""

	if len(args) >= 2 {
		oldName = args[0]
		newName = args[1]
	} else if len(args) >= 1 {
		newName = args[0]
	}

	if oldName != "" && newName != "" {
		message := GitStr("remote", "rename", oldName, newName)
		error := strings.Contains(message, "fatal:")

		if error {
			fmt.Println(message)
		} else {
			fmt.Println(message)

			Git("remote", "-v")
		}
	}
}

func show() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	Git("remote", "show", name)
}

func prune() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	Git("remote", "prune", name)
}

func url() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	Git("remote", "get-url", name)
}

func list() {
	message := GitStr("remote", "-v")

	if len(message) > 0 {
		fmt.Println()
		fmt.Println("Registered remote repositories:")
		fmt.Println()
		fmt.Print(message)
	} else {
		fmt.Printf(`
Remote repositories are empty.

Use the following command to add a repository:
    
    %[1]s repo <url>
    %[1]s repo <name> <url>
    %[1]s repo + <url>
    %[1]s repo + <name> <url>

Examples:

    %[1]s repo https://github.com/pytorch/pytorch
    %[1]s repo origin https://github.com/pytorch/pytorch
    %[1]s repo + https://github.com/pytorch/pytorch
    %[1]s repo + origin https://github.com/pytorch/pytorch
    `, Command)
	}
}
