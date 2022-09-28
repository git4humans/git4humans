package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Save() {
	args := os.Args[2:]
	params := remove(args, "-m")
	message := ""

	status := GitStr("status")
	noCommit := strings.Contains(status, "nothing to commit")

	if noCommit {
		fmt.Println(status)
		return
	}

	if len(params) > 0 {
		message = params[0]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nMessage: ")

		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			message = strings.Replace(string(input), "\n", "", -1)
		}
	}

	if len(message) > 0 {
		if NotGit() {
			Git("init")
		}

		if HasUpdate() {
			Git("add", ".")
		}

		Git("commit", "-m", message)
	} else {
		fmt.Println("\nUnable to save a commit without message.")
	}
}
