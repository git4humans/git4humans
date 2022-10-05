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
	command := os.Args[1]

	if command == "rename" {
		RenameHelp()
	} else {
		fmt.Printf(`
Move a file, a directory, or a symlink into another location.

Usage: 

    %[1]s move <file> <target>
    %[1]s mv <file> <target>
		`, Command)
	}
}

func RenameHelp() {
	fmt.Printf(`
Rename a file, a directory, or a symlink.

Usage: 

    %[1]s rename <oldname> <newname>
    %[1]s ren <oldname> <newname>
	`, Command)
}
