package git4humans

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Git(subcommand string, args ...string) {
	fmt.Println(GitStr(subcommand, args...))
}

func GitStr(subcommand string, args ...string) string {
	params := append([]string{subcommand}, args...)
	cmd := exec.Command("git", params...)

	var out bytes.Buffer
	var err bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &err

	if cmd.Run() != nil {
		error := err.String()
		notGit := strings.Contains(error, "not a git repository")

		if notGit {
			error = fmt.Sprintf(`
%[1]s
Use the following command to create a new Git repository:

    %[2]s new 


Note:
The '%[2]s new' will initialize a Git repository, then add all files into staging, and do initial commit.

Use '%[2]s init' if you only want to initialize.
Use '%[2]s +' if you only want to initialize and add files into staging.
Use '%[2]s save' if you only want to initialize, add all files, and commit with your own message.`, error, Command)
		}

		return error
	} else {
		return out.String()
	}
}

func NotGit() bool {
	status := GitStr("status")
	result := strings.Contains(status, "not a git repository")

	return result
}

func HasUpdate() bool {
	status := GitStr("status")

	untracked := strings.Contains(status, "Untracked files:")
	unstaged := strings.Contains(status, "Changes not staged for commit:")

	return untracked || unstaged
}
