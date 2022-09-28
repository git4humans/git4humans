package git4humans

import (
	"fmt"
	"os"
)

func User() {
	args := os.Args[2:]
	global := contains(args, "--global")
	params := remove(args, "--global")

	if len(params) >= 2 {
		name := params[0]
		email := params[1]

		if global {
			Git("config", "--global", "user.name", name)
			Git("config", "--global", "user.email", email)
		} else {
			Git("config", "user.name", name)
			Git("config", "user.email", email)
		}

		fmt.Println()
	}

	if global {
		name := GitStr("config", "--global", "user.name")
		email := GitStr("config", "--global", "user.email")

		fmt.Printf(`
Git user configuration (global):

user.name   %[1]suser.email  %[2]s
        `, name, email)
	} else {
		name := GitStr("config", "user.name")
		email := GitStr("config", "user.email")

		fmt.Printf(`
Git user configuration (local):

user.name   %[1]suser.email  %[2]s
        `, name, email)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func remove(s []string, str string) []string {
	for i, v := range s {
		if v == str {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
