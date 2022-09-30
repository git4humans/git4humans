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
		StartUsage()
		return
	}

	fmt.Print(`
This will create a Git repository in your working directory, then add all files into staging  and do an initial commit:

    git init
    git add .
    git commit -m "Initial commit"

`)
	fmt.Print("Continue? (y/n) ")

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

func StartUsage() {
	fmt.Println()
}
