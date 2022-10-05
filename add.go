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
		AddHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	if NotGit() {
		fmt.Println()
		fmt.Println("Warn: your project is not a Git repository.")
		fmt.Println()
		fmt.Print("Create a Git repository (Y/N)? ")

		reader := bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadLine()

		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
			confirm = strings.ToLower(confirm)
		}

		if confirm == "y" {
			fmt.Println()
			Git("init")
		}
	}

	if NotGit() {
		fmt.Println()
		fmt.Println("Abort: not a Git repository, cannot stage your file(s).")
		return
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
    
    %[1]s %[2]s file.txt
    %[1]s %[2]s file1.txt file2.txt file3.txt
    %[1]s %[2]s file1.txt dir2/file2.txt
    %[1]s %[2]s file.txt dir/
    %[1]s %[2]s --all 
    %[1]s %[2]s .
            `, Command, command)
		}
	}
}

func AddHelp() {
	command := os.Args[1]

	fmt.Printf(`
Adds new or modified file(s) to the staging area for inclusion in the next commit.

Usage: 

    %[1]s + [<file>]
    %[1]s a [<file>]
    %[1]s add [<file>]

Examples:

    %[1]s %[2]s file.txt
    %[1]s %[2]s file1.txt file2.txt file3.txt
    %[1]s %[2]s file1.txt dir2/file2.txt
    %[1]s %[2]s file.txt dir/
    %[1]s %[2]s --all
    %[1]s %[2]s .
	`, Command, command)
}
