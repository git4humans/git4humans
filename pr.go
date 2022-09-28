package git4humans

import "os"

func Pr() {
	args := os.Args[2:]

	Git("request-pull", args...)
}
