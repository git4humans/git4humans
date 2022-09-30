package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Save() {
	if IsHelp() {
		SaveUsage()
		return
	}

	if NotGit() {
		fmt.Println()
		fmt.Println("Warn: your project is not a Git repository.")
		fmt.Println()
		fmt.Print("Create a Git repository? (y/n) ")

		reader := bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadLine()

		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
		}

		yes := confirm == "Y" || confirm == "y"

		if yes {
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
				fmt.Println()

				Git("status")

				fmt.Println()
				fmt.Println("Warn: this will save all changes in your branch.")
			}

			fmt.Println()
			fmt.Print("Save with message: ")

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
			fmt.Println("Abort: cannot save without a message.")
		}
	} else {
		fmt.Println()
		fmt.Println("Abort: no changes to save in your branch.")
	}
}

func SaveUsage() {
	fmt.Println()
}
