package git4humans

import (
	"fmt"
	"os"
)

func LastLog() {
	if IsHelp() {
		LastHelp()
		return
	}

	args := os.Args[2:]
	count := "1"

	if len(args) > 0 {
		count = args[0]
	}

	Git("log", "-"+count, "HEAD")
}

func LastHelp() {
	fmt.Printf(`
Show the last N commit in the log history.

Usage: %[1]s last <N>

The argument N specifies the number of last commits that you want to show. For example, '%[1]s last 2' command will show the last 2 commits in your history. If you don't specify, it will show only the one last commit.
		`, Command)
}
