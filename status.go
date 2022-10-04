package git4humans

import "os"

func Status() {
	args := os.Args[2:]

	Git("status", args...)
}

func ShortStatus() {
	Git("status", "-s")
}
