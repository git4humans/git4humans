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

	params := args
	params = remove("-m", params)
	params = remove("--message", params)

	if HasUpdate() || HasCommit() {
		if len(params) > 0 {
			message = params[0]
		} else {
			Git("status")
			fmt.Println()

			fmt.Println("This will save all changes in your branch.")
			fmt.Print("\nSave with message: ")

			reader := bufio.NewReader(os.Stdin)
			input, _, _ := reader.ReadLine()

			if len(input) > 0 {
				message = strings.Replace(string(input), "\n", "", -1)
				message = strings.Trim(message, "\n")
			}
		}

		if len(message) > 0 {
			fmt.Println()

			if NotGit() {
				Git("init")
			}

			if HasUpdate() {
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
