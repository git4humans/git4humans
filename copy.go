package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Clone a remote repository
func Copy() {
	if IsHelp() {
		CopyHelp()
		return
	}

	args := os.Args[2:]

	if len(args) > 0 {
		Git("clone", args...)
	} else {
		url := ""
		dir := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("\nRepository URL: ")

		inputUrl, _, _ := reader.ReadLine()

		if len(inputUrl) > 0 {
			url = strings.Replace(string(inputUrl), "\n", "", -1)
		}

		fmt.Print("Target directory: ")

		inputDir, _, _ := reader.ReadLine()

		if len(inputDir) > 0 {
			dir = strings.Replace(string(inputDir), "\n", "", -1)
		}

		fmt.Println("")

		if len(url) > 0 {
			if len(dir) > 0 {
				Git("clone", url, dir)
			} else {
				Git("clone", url)
			}
		}
	}
}

func CopyHelp() {
	fmt.Printf(`
Copy an existing repository into a new directory.

usage: %[1]s copy <url> <directory>

Using the command without specifying a target directory will copy the repository into a new directory with the name specified in the repository URL.

examples:
    %[1]s copy https://github.com/pytorch/pytorch
    %[1]s copy https://github.com/pytorch/pytorch torch
	`, Command)
}
