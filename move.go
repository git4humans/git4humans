package git4humans

import (
	"fmt"
	"os"
)

// Move or rename a file
func Move() {
	if IsHelp() {
		MoveHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	if len(args) >= 2 {
		oldName := args[0]
		newName := args[1]
		options := args[2:]

		if command == "rename" {
			Git("mv", append([]string{oldName, newName}, options...)...)
		} else {
			Git("mv", oldName, newName)
		}
	}
}

func MoveHelp() {
	fmt.Printf(``)
}

func RenameHelp() {
	fmt.Printf(``)
}
