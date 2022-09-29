package main

import "github.com/git4humans/git4humans"

func main() {
	git4humans.Command = "gh"
	git4humans.Handle()
}
