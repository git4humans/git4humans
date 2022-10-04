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
Warn: this will create a Git repository in your project, then automatically stage all files and do an initial commit.

Equivalent of:

    git init
    git add .
    git commit -m "Initial commit"

`)
	fmt.Print("Press enter to continue, c to cancel: ")

	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	confirm := ""

	if len(input) > 0 {
		confirm = strings.Replace(string(input), "\n", "", -1)
	}

	if confirm != "c" {
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

The command creates a Git repository in your current working directory, then automatically stage all files and do an initial commit. 

Equivalent of:

    git init
    git add .
    git commit -m "Initial commit"
	`, Command)
}
