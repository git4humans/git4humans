package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Clone a remote repository
func Copy() {
	args := os.Args[2:]

	if len(args) > 0 {
		Git("clone", args...)
	} else {
		url := ""
		dir := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("\nRepository URL: ")

		textUrl, _, _ := reader.ReadLine()

		if len(textUrl) > 0 {
			url = strings.Replace(string(textUrl), "\n", "", -1)
		}

		fmt.Print("Target directory: ")

		textDir, _, _ := reader.ReadLine()

		if len(textDir) > 0 {
			dir = strings.Replace(string(textDir), "\n", "", -1)
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
