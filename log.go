package git4humans

import "os"

func Log() {
	args := os.Args[2:]

	Git("log", args...)
}
