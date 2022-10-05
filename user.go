package git4humans

import (
	"fmt"
	"os"
)

func User() {
	if IsHelp() {
		UserHelp()
		return
	}

	args := os.Args[2:]
	global := HasOption("--global", args) || HasOption("-g", args)

	if global {
		args = RemoveOption("--global", args)
		args = RemoveOption("-g", args)
	}

	if len(args) >= 2 {
		name := args[0]
		email := args[1]

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

func UserHelp() {
	command := os.Args[1]

	fmt.Printf(`
Show or update user Git configuration (local or global).

Usage:

    %[1]s user <option> <name> <email>
    %[1]s u <option> <name> <email>

Using the command without --global or -g option will show or update user configuration for the local current working Git repository.

Examples:

    %[1]s %[2]s 
    %[1]s %[2]s --global 
    %[1]s %[2]s -g 
    
    %[1]s %[2]s "Salman S" "salkuadrat@gmail.com" 
    %[1]s %[2]s --global "Salman S" "salkuadrat@gmail.com"
    %[1]s %[2]s -g "Salman S" "salkuadrat@gmail.com"
	`, Command, command)
}
