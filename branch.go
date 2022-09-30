package git4humans

import (
	"fmt"
	"os"
	"strings"
)

func Branch() {
	if IsHelp() {
		BranchUsage()
		return
	}

	args := os.Args[2:]

	if len(args) > 0 {
		branch := args[0]

		fmt.Println("Creating branch " + branch + "...")
		Git("branch", branch)

		fmt.Println("Switching to " + branch + "...")
		Git("switch", branch)

		showBranch()
	} else {
		showBranch()
	}
}

func SwitchBranch() {
	if IsHelp() {
		SwitchUsage()
		return
	}

	args := os.Args[2:]

	message := GitStr("switch", args...)
	error := strings.Contains(message, "fatal:")

	if error {
		fmt.Println(message)
	} else {
		showBranch()
	}
}

func showBranch() {
	branch := GitStr("branch", "--show-current")

	fmt.Println("On branch " + branch)
	fmt.Println()
	fmt.Println("List branch:")

	Git("branch")
}

func BranchUsage() {
	fmt.Println()
}

func SwitchUsage() {
	fmt.Println()
}
