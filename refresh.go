package git4humans

import "fmt"

func Refresh() {
	if IsHelp() {
		RefreshHelp()
	} else {
		Git("init")
	}
}

func RefreshHelp() {
	fmt.Printf(`
Refresh an existing Git repository in the current working directory.

usage: %[1]s refresh
	`, Command)
}
