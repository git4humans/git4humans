package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Create a working directory for a new project.
// Then intialize a fresh Git repository into the working directory.
//
// mdkir <project>
// cd <project>
// git init
func New() {
	if IsHelp() {
		NewHelp()
		return
	}

	args := os.Args[2:]
	dir := ""

	if len(args) > 0 {
		dir = args[0]
	}

	if len(dir) <= 0 {
		fmt.Println()
		fmt.Print("Directory: ")

		reader := bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			dir = strings.Replace(string(input), "\n", "", -1)
		}
	}

	if len(dir) <= 0 {
		fmt.Println()
		fmt.Println("Abort: you should specify directory for the new Git repository.")
		fmt.Println()
		fmt.Printf("Use '%[1]s start' to create a fresh Git in this current working directory.", Command)
		fmt.Println()
		return
	}

	err := os.Mkdir(dir, os.ModePerm)

	if err != nil {
		fmt.Println("Abort: failed creating directory for the new Git repository.")
		return
	}

	err = os.Chdir(dir)

	if err != nil {
		panic(err)
	} else {
		Git("init")
		os.OpenFile(".gitignore", os.O_RDONLY|os.O_CREATE, 0666)
	}
}

func NewHelp() {
	fmt.Printf(`
Create a new Git repository in a new directory.

Usage: 

    %[1]s new <directory>
    %[1]s n <directory>

The command creates a new directory (mkdir), then creates a Git repository in the new directory (basically a .git directory with subdirectories for objects, refs/heads, refs/tags, and template files) and adds an empty .gitignore file into the new directory.
	`, Command)
}
