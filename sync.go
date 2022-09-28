package git4humans

import "os"

func Sync() {
	args := os.Args[2:]

	Git("pull", args...)
}
