package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Add untracked files into staging.
// Automatically initialize if the repo is not initialized.
func Add() {
	if IsHelp() {
		AddUsage()
		return
	}

	args := os.Args[2:]

	if NotGit() {
		Git("init")
	}

	if len(args) > 0 {
		Git("add", args...)
		Git("status")
	} else {
		files := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("\nFile(s) to add: ")

		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			files = strings.Replace(string(input), "\n", "", -1)
		}

		if len(files) > 0 {
			Git("add", strings.Fields(files)...)
			Git("status")
		} else {
			fmt.Printf(`
You should specify file(s) to add.

Examples:
    
    %[1]s + file.txt
    %[1]s + file1.txt file2.txt file3.txt
    %[1]s + file1.txt dir2/file2.txt
    %[1]s + file.txt dir/
    %[1]s + --all 
    %[1]s + .
            `, Command)
		}
	}
}

func AddUsage() {
	fmt.Printf(`
Add file contents to the index

usage: %[1]s + [<file>]

examples:
    %[1]s + file.txt
    %[1]s + file1.txt file2.txt file3.txt
    %[1]s + file1.txt dir2/file2.txt
    %[1]s + file.txt dir/
    %[1]s + --all
    %[1]s + .
	`, Command)
}
