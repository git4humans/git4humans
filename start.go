package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Start a new git repository:
//
// git init
// git add .
// git commit -m "Initial commit"
func Start() {
	if IsHelp() {
		StartHelp()
		return
	}

	fmt.Print(`
Warn: this will create a fresh Git repository in your project, then automatically stage all files and do initial commit, equivalent as the following:

    git init
    git add .
    git commit -m "Initial commit"

`)
	fmt.Print("Do you want to continue? (y/n) ")

	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	confirm := ""

	if len(input) > 0 {
		confirm = strings.Replace(string(input), "\n", "", -1)
	}

	yes := confirm == "Y" || confirm == "y"

	if yes {
		args := os.Args[2:]

		Git("init", args...)

		if HasUpdate() {
			Git("add", ".")
			Git("commit", "-m", "Initial commit")
		}
	}
}

func StartHelp() {
	fmt.Printf(`
Start a Git repository in your project.

usage: %[1]s start

The command creates a fresh Git repository in this current working directory, then automatically stage all files and do initial commit, equivalent as:

    git init
    git add .
    git commit -m "Initial commit"
	`, Command)
}
