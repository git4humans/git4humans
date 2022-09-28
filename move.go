package git4humans

import "os"

// Move or rename a file
func Move() {
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
