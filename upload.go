package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Upload changes into remote repository
// Another way of git push
func Upload() {
	if IsHelp() {
		UploadHelp()
		return
	}

	reader := bufio.NewReader(os.Stdin)

	if NotGit() {
		fmt.Println()
		fmt.Println("Warn: this project is not a Git repository.")
		fmt.Println()
		fmt.Print("Create a Git repository (Y/N)? ")

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

	if NoRepo(remote) {
		fmt.Println()
		fmt.Printf("Warn: remote repository %[1]s is not found.", remote)
		fmt.Println()
		fmt.Println()
		fmt.Printf("Add URL for remote %[1]s: ", remote)

		url := ""
		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			url = strings.Replace(string(input), "\n", "", -1)
		}

		if len(url) > 0 {
			fmt.Println()
			Git("remote", "add", remote, url)
		}
	}

	if NoRepo(remote) {
		fmt.Println()
		fmt.Println("Abort: you should add remote repository to publish.")
		return
	}

	if len(args) < 2 {
		branch := GitStr("branch", "--show-current")
		branch = strings.Trim(branch, "\n")
		url := GitStr("remote", "get-url", remote)
		url = strings.Trim(url, "\n")

		fmt.Println("On branch " + branch)
		fmt.Println()

		fmt.Println("Warn:")
		fmt.Printf(`This will upload your branch %[1]s to the branch %[2]s of remote %[3]s (%[4]s)
		`, localBranch, remoteBranch, remote, url)
		fmt.Println()
		fmt.Print("Press enter to continue, c to cancel: ")

		input, _, _ := reader.ReadLine()
		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
			confirm = strings.ToLower(confirm)
		}

		if confirm == "c" {
			return
		}
	}

	if HasUpdate() {
		fmt.Println()
		fmt.Println("You have some unstaged changes in your project:")
		Git("status", "-s")

		fmt.Println()
		fmt.Print("Save the changes (Y/N)? ")

		input, _, _ := reader.ReadLine()
		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
			confirm = strings.ToLower(confirm)
		}

		if confirm == "y" {
			fmt.Println()
			fmt.Print("Saving changes with message: ")

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
				fmt.Println("Cannot save changes without a message.")
			}
		}
	}

	if CanUpload() {
		fmt.Println()
		fmt.Printf("Uploading %[1]s to %[2]s/%[3]s...", localBranch, remote, remoteBranch)
		fmt.Println()

		// Git("push", name, branch)
		// Use a standard command to print results
		cmd := exec.Command("git", "push", remote, remoteBranch)
		response, _ := cmd.CombinedOutput()

		fmt.Println()
		fmt.Print(string(response))
	} else {
		fmt.Println()
		fmt.Printf("Your branch has nothing to upload. See %[1]s status.", Command)
		fmt.Println()
	}
}

func CanUpload() bool {
	status := GitStr("status")
	canPublish := strings.Contains(status, "Your branch is ahead of")

	return canPublish
}

func UploadHelp() {
	fmt.Printf(``)
}
