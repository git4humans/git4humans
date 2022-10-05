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

Usage: 

    %[1]s refresh
    %[1]s r
	`, Command)
}
