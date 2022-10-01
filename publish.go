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
		PublishHelp()
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
	remote := "origin"
	remoteBranch := ""
	localBranch := strings.Trim(GitStr("branch", "--show-current"), "\n")

	if len(args) >= 2 {
		remote = args[0]
		remoteBranch = args[1]
	} else {
		remoteBranch = localBranch
	}

	if len(args) < 2 {
		url := ""

		if HasRepo(remote) {
			url = GitStr("remote", "get-url", remote)
			url = strings.Trim(url, "\n")
			url = "(" + url + ")"
		}

		fmt.Printf(`
Warn: this will publish your local branch %[1]s to the branch %[2]s of remote %[3]s %[4]s 
		`, localBranch, remoteBranch, remote, url)
		fmt.Println()
		fmt.Print("Continue? (y/n) ")

		input, _, _ := reader.ReadLine()
		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
		}

		abort := confirm != "Y" && confirm != "y"

		if abort {
			return
		}
	}

	if NoRepo(remote) {
		fmt.Println()
		fmt.Println("Remote repository " + remote + " is not found.")
		fmt.Println()
		fmt.Print("Add URL for " + remote + ": ")

		url := ""
		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			url = strings.Replace(string(input), "\n", "", -1)
		}

		if len(url) > 0 {
			Git("remote", "add", remote, url)
		}
	}

	if NoRepo(remote) {
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
		fmt.Print("Save this changes? (y/n) ")

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
		fmt.Printf("Publishing from local %[1]s to remote %[2]s %[3]s...", localBranch, remote, remoteBranch)
		fmt.Println()

		// Git("push", name, branch)
		// Use a standard command to print results
		cmd := exec.Command("git", "push", remote, remoteBranch)
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

func PublishHelp() {
	fmt.Printf(``)
}
