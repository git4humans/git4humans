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
	fmt.Print(`
This will create a Git repository in your current working directory, then add all files into staging and do initial commit:

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
	} else {
		fmt.Println()
		fmt.Println("Abort: start command is cancelled.")
	}
}
