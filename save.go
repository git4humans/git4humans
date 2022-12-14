package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Save() {
	if IsHelp() {
		SaveHelp()
		return
	}

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
		fmt.Println("Abort: not a Git repository, cannot save.")
		return
	}

	args := os.Args[2:]
	message := ""

	if HasFlag("-m", args) {
		message = GetFlag("-m", args)
		args = RemoveFlag("-m", args)
	}

	if HasFlag("--message", args) {
		message = GetFlag("--message", args)
		args = RemoveFlag("--message", args)
	}

	if HasUpdate() || HasCommit() {
		if len(message) <= 0 {
			// only show details and warning
			// if the branch has some changes
			// and no additional arguments (containing the list of files to save)
			if HasUpdate() && len(args) <= 0 {
				ListChanges()
			}

			fmt.Println()
			fmt.Print("Saving changes with message: ")

			reader := bufio.NewReader(os.Stdin)
			input, _, _ := reader.ReadLine()

			if len(input) > 0 {
				message = strings.Replace(string(input), "\n", "", -1)
				message = strings.Trim(message, "\n")
			}
		}

		if len(message) > 0 {
			fmt.Println()

			if len(args) > 0 {
				Git("add", args...)
			} else if HasUpdate() {
				Git("add", ".")
			}

			if HasCommit() {
				Git("commit", "-m", message)
			}
		} else {
			fmt.Println()
			fmt.Println("Abort: cannot save changes without a message.")
		}
	} else {
		fmt.Println()
		fmt.Println("Abort: no changes to save in your branch.")
	}
}

func SaveHelp() {
	command := os.Args[1]

	fmt.Printf(`
Record changes in your current local repository.

Usage: 

    %[1]s save [<file>] <option>
    %[1]s sv [<file>] <option>

The command stages the file(s) into the staging area to be ready for commit, then commit with the specified message. If you use the command without -m option, it will ask to type the message.

Examples: 

    %[1]s %[2]s  
    %[1]s %[2]s . 
    %[1]s %[2]s -m "Commit message"
    %[1]s %[2]s file1.txt file2.txt 
    %[1]s %[2]s file1.txt -m "Describe the usage"
	`, Command, command)
}
