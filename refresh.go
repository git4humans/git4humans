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
Refresh the existing Git repository in your project.

usage: %[1]s refresh
	`, Command)
}
