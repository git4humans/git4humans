package git4humans

import (
	"fmt"
	"os"
)

func User() {
	if IsHelp() {
		UserUsage()
		return
	}

	args := os.Args[2:]
	params := args

	global := contains("--global", args) || contains("-g", args)

	if global {
		params = remove("--global", params)
		params = remove("-g", params)
	}

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

func UserUsage() {
	fmt.Println()
}
