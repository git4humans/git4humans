package main

import "github.com/git4humans/git4humans"

func main() {
	git4humans.Command = "gt"
	git4humans.Handle()
}
