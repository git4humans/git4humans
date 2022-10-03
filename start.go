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
Warn: this will create a fresh Git repository in your project, then automatically stage all files and do initial commit.

An equivalent of:

    git init
    git add .
    git commit -m "Initial commit"

`)
	fmt.Print("Enter to continue, q to quit: ")

	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	confirm := ""

	if len(input) > 0 {
		confirm = strings.Replace(string(input), "\n", "", -1)
	}

	if confirm != "q" {
		args := os.Args[2:]

		fmt.Println()

		Git("init", args...)

		if HasUpdate() {
			Git("add", ".")
			Git("commit", "-m", "Initial commit")
		}
	}
}

func StartHelp() {
	fmt.Printf(`
Start a Git repository for your project.

usage: %[1]s start

The command creates a fresh Git repository in your current working directory, then automatically stage all files and do initial commit. 

An equivalent of:

    git init
    git add .
    git commit -m "Initial commit"
	`, Command)
}
