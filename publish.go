package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Publish changes into remote repository
// Another way of git push
func Publish() {
	if IsHelp() {
		PublishUsage()
		return
	}

	reader := bufio.NewReader(os.Stdin)

	if NotGit() {
		fmt.Println()
		fmt.Println("Warn: this project is not a Git repository.")
		fmt.Println()
		fmt.Print("Create a Git repository? (y/n) ")

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
		fmt.Println("Abort: not a Git repository, cannot publish.")
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
		branch = strings.Trim(branch, "\n")
	}

	if NoRepo(name) {
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

	if NoRepo(name) {
		fmt.Println("Abort: you should add remote repository to publish.")
		return
	}

	if HasUpdate() {
		fmt.Println()

		Git("status")

		fmt.Println()
		fmt.Println()
		fmt.Println("Warn: you have some unstaged changes in this branch.")
		fmt.Println()
		fmt.Print("Want to save this changes? (y/n) ")

		input, _, _ := reader.ReadLine()
		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
		}

		yes := confirm == "Y" || confirm == "y"

		if yes {
			fmt.Println()
			fmt.Print("Save with message: ")

			input, _, _ := reader.ReadLine()
			message := ""

			if len(input) > 0 {
				message = strings.Replace(string(input), "\n", "", -1)
			}

			if len(message) > 0 {
				fmt.Println()

				Git("add", ".")
				Git("commit", "-m", message)
			} else {
				fmt.Println()
				fmt.Println("Cannot save without a message.")
			}
		}
	}

	if CanPublish() {
		fmt.Println()
		fmt.Printf("Publishing into %[1]s %[2]s...", name, branch)
		fmt.Println()

		// Git("push", name, branch)
		// Use a standard command to print results
		cmd := exec.Command("git", "push", name, branch)
		response, _ := cmd.CombinedOutput()

		fmt.Println()
		fmt.Print(string(response))
	} else {
		fmt.Println()
		fmt.Printf("Your branch has nothing to publish. See %[1]s status.", Command)
		fmt.Println()
	}
}

func CanPublish() bool {
	status := GitStr("status")
	canPublish := strings.Contains(status, "Your branch is ahead of")

	return canPublish
}

func PublishUsage() {
	fmt.Println()
}
