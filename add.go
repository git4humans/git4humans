package git4humans

import "os"

// Add untracked files into staging.
// Automatically initialize if the repo is not initialized.
func Add() {
	args := os.Args[2:]

	if NotGit() {
		Git("init")
	}

	if len(args) > 0 {
		Git("add", args...)
	} else {
		Git("add", ".")
	}

	Git("status")
}
