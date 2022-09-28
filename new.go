package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Initializing a new git repository:
//
// git init
// git add .
// git commit
func New() {
	fmt.Print(`
This action will initialize a new Git repository, then add all files into staging and do initial commit:

    git init 
    git add .
    git commit -m "Initial commit"

`)
	fmt.Print("Continue? (y/n) ")

	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()

	if len(input) > 0 {
		confirm := strings.Replace(string(input), "\n", "", -1)
		yes := confirm == "Y" || confirm == "y"

		if yes {
			args := os.Args[2:]

			Git("init", args...)
			Git("add", ".")
			Git("commit", "-m", "Initial commit")
		}
	}
}
