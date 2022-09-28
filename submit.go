package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// git push
func Submit() {
	reader := bufio.NewReader(os.Stdin)

	if NotGit() {
		fmt.Println()
		fmt.Println("This is not a Git repository.")
		fmt.Println()
		fmt.Print("Do you want to init a Git repository? (y/n) ")

		confirm := ""
		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
		}

		yes := confirm == "Y" || confirm == "y"

		if yes {
			Git("init")
			Git("add", ".")
			Git("commit", "-m", "Initial commit")
		}
	}

	if NotGit() {
		fmt.Println("Abort: Cannot submit a non Git repository.")
		return
	}

	args := os.Args[2:]
	name := "origin"
	branch := ""

	if len(args) >= 2 {
		name = args[0]
		branch = args[1]
	} else {
		branch = GitStr("branch", "--show-current")
	}

	if !HasRepo(name) {
		fmt.Println()
		fmt.Println("Remote repository " + name + " is not found.")
		fmt.Println()
		fmt.Print("Add URL for " + name + ": ")

		url := ""
		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			url = strings.Replace(string(input), "\n", "", -1)
		}

		if len(url) > 0 {
			Git("remote", "add", name, url)
		}
	}

	if !HasRepo(name) {
		fmt.Println("Abort: Cannot submit to an undefined remote repository.")
		return
	}

	if HasUpdate() {
		fmt.Println()
		fmt.Println("There are some changes in your project.")
		fmt.Println()
		fmt.Print("Do you want to commit? (y/n) ")

		input, _, _ := reader.ReadLine()
		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
		}

		yes := confirm == "Y" || confirm == "y"

		if yes {
			fmt.Print("Message: ")

			input, _, _ := reader.ReadLine()
			message := ""

			if len(input) > 0 {
				message = strings.Replace(string(input), "\n", "", -1)
			}

			if len(message) > 0 {
				Git("commit", "-m", message)
			} else {
				fmt.Println("Cannot commit without a message.")
				fmt.Println()
			}
		}
	}

	if NoCommit() {
		Git("status")
		return
	}

	Git("push", name, branch)
}
