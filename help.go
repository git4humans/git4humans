package git4humans

import "fmt"

func Help() {
	fmt.Printf(`
Git for Humans version %[1]s

usage: %[2]s [<options>] <command> [<args]

These are common commands you will most likely use in everyday situation.

Start a working area
    copy  	Copy an existing repository to a new directory
    new		Create a Git repository in a new directory 
    start	Start a fresh Git repository in the current directory
    refresh     Refresh or reinitialize the existing Git repository in the current directory

Work on the current change 
    + 		Add file contents to the staging area 
    - 		Delete file(s) from the directory and the staging area
    move 	Move a file, a directory, or a symlink 
    rename 	Rename a file, a directory, or a symlink 
    restore	Restore deleted file(s)
    save 	Record changes to the repository (staging and committing file(s))

Collaborate with others
    repo 	Manage remote repositories (e.g., list, add, delete)
    fetch 	Fetch everything from a remote repository
    sync 	Fetch from and integrate with a remote repository or a local branch 
    publish     Publish changes to a remote repository (shortcut: '%[2]s pub')

Grow, mark, and tweak your common history 
    branch 	List, create, or delete branches 
    switch 	Switch to another branch 
    merge 	Join two or more branches 
    reset 	Reset current HEAD to a specified state
    tag 	Create, list, delete, or verify a tag object

Examine history and state 
    status 	Show current status of the branch (shortcut: '%[2]s s') 
    log 	Show commit logs (shortcut: '%[2]s l')
    diff 	Show changes between commits, commit and working tree, etc.

Configurations 
    config	Show or update configuration
    user	Show or update user configuration (local or global)

You may use other original commands from Git.
Ex: you can use '%[2]s push origin main' and 'git push origin main' interchangeably.
	`, Version, Command)
}
