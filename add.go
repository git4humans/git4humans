package git4humans

import (
	"fmt"
	"os"
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
	} else {
		Git("add", ".")
	}

	Git("status")
}

func AddUsage() {
	fmt.Printf(`
Add file contents to the index

usage: %[1]s + [<file>]

examples:
    %[1]s + file.txt
    %[1]s + file1.txt file2.txt
    %[1]s + file1.txt dir2/file2.txt
	`, Command)
}
