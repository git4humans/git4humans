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

	args := os.Args[2:]
	message := ""

	if HasFlag("-m", args) {
		message = GetFlag("-m", args)
		args = RemoveFlag("-m", args)
	} else if HasFlag("--message", args) {
		message = GetFlag("--message", args)
		args = RemoveFlag("--message", args)
	}

	if NotGit() {
		Git("init")
	}

	if HasUpdate() || HasCommit() {
		if len(message) <= 0 {
			if len(args) <= 0 && HasUpdate() {
				Git("status")

				fmt.Println()
				fmt.Println("This will save all changes in your branch.")
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
		fmt.Println("Abort: no changes to save in your project.")
	}
}

func SaveUsage() {
	fmt.Println()
}
