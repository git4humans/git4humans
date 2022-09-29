package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Save() {
	args := os.Args[2:]
	message := ""

	params := args
	params = remove(params, "-m")
	params = remove(params, "--message")

	Git("status")

	if HasCommit() {
		if len(params) > 0 {
			message = params[0]
		} else {
			fmt.Print("\nMessage: ")

			reader := bufio.NewReader(os.Stdin)
			input, _, _ := reader.ReadLine()

			if len(input) > 0 {
				message = strings.Replace(string(input), "\n", "", -1)
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

			Git("commit", "-m", message)
		} else {
			fmt.Println()
			fmt.Println("Abort: cannot commit without a message.")
		}
	}
}

func HasCommit() bool {
	return !NoCommit()
}

func NoCommit() bool {
	status := GitStr("status")
	noCommit := strings.Contains(status, "nothing to commit")

	return noCommit
}
