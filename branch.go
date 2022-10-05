package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Branch() {
	if IsHelp() {
		BranchHelp()
		return
	}

	args := os.Args[2:]

	if len(args) > 0 {
		branch := args[0]

		fmt.Println("Creating branch " + branch + "...")
		Git("branch", branch)

		fmt.Println()
		fmt.Printf("Want to switch to the new branch %[1]s (Y/N)? ", branch)

		reader := bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadLine()

		confirm := ""

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
			confirm = strings.ToLower(confirm)
		}

		if confirm == "y" {
			GitStr("switch", branch)
		}

		fmt.Println()
	}

	showBranch()
}

func CurrentBranch() string {
	branch := GitStr("branch", "--show-current")
	branch = strings.TrimRight(branch, "\n")

	return branch
}

func SwitchBranch() {
	if IsHelp() {
		SwitchHelp()
		return
	}

	args := os.Args[2:]

	message := GitStr("switch", args...)
	error := strings.Contains(message, "fatal:")

	if error {
		fmt.Println()

		if len(args) > 0 {
			branch := args[0]
			notFound := strings.Contains(message, "invalid reference: "+branch)

			if notFound {
				fmt.Println("Branch with name " + branch + " is not found.")
			} else {
				fmt.Println(message)
			}
		} else {
			fmt.Println("You should specify the name of branch to switch.")
		}

	} else {
		showBranch()
	}
}

func showBranch() {
	if isGit() {
		fmt.Println("On branch " + CurrentBranch())
		fmt.Println()
		fmt.Println("List branch:")
	}

	Git("branch")
}

func BranchHelp() {
	command := os.Args[1]

	fmt.Printf(`
Manage branches of your project (list, create, delete, etc.)

Usage:

    %[1]s branch
    %[1]s br

    %[1]s branch <new-branch>
    %[1]s br <new-branch>

    %[1]s branch delete <branch>
    %[1]s branch del <branch>
    %[1]s branch - <branch>

    %[1]s br delete <branch>
    %[1]s br del <branch>
    %[1]s br - <branch>

Examples 

Show the list of existing branch: 

    %[1]s %[2]s

Add a new branch: 

    %[1]s %[2]s test

Delete an existing branch:

    %[1]s %[2]s delete test
    %[1]s %[2]s d test
    %[1]s %[2]s - test
	`, Command, command)
}

func SwitchHelp() {
	command := os.Args[1]

	fmt.Printf(`
Switch from your current branch to another existing branch.

Usage:

    %[1]s switch <branch>
    %[1]s sw <branch>

Examples:

    %[1]s switch test 
    %[1]s sw test
	`, Command, command)
}
