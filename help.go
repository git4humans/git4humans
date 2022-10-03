package git4humans

import "fmt"

func Help() {
	fmt.Printf(`
Git for Humans version %[1]s

usage: %[2]s [<options>] <command> [<args]

These are common commands you will most likely use in everyday situation.

Start a working area
    copy  	Copy an existing Git repository to a new directory
    new		Create a new Git repository in a new directory 
    start	Start a new Git repository in the current working directory
    refresh     Refresh the existing Git repository in the current working directory

Work on the current change 
    + 		Add file(s) to the staging area for inclusion in the next commit
    - 		Delete file(s), then stages the deletion for the next commit
    move 	Move a file, a directory, or a symlink 
    rename 	Rename a file, a directory, or a symlink 
    restore	Restore deleted file(s)
    save 	Record changes in your local repository

Collaborate with others
    repo 	Manage remote repositories (list, add, delete, etc.)
    fetch 	Fetch everything from a remote repository
    sync 	Fetch from and integrate with a repository or a local branch 
    publish     Publish changes to a remote repository (shortcut: '%[2]s pub')

Grow, mark, and tweak your common history 
    branch 	Manage branches (list, create, delete, etc.) 
    switch 	Switch from your current branch to another branch 
    merge 	Join two or more branches
    reset 	Reset current HEAD to a specified state
    tag 	Create, list, delete, or verify a tag object

Examine history and state 
    status 	Show the status of the working copy (shortcut: '%[2]s s') 
    log 	Show a reversed-ordered list of commits (shortcut: '%[2]s l')
    diff 	Show changes between commits, commit and working tree, etc.

Configurations 
    config	Show or update configuration
    user	Show or update user configuration (local or global)

You also can use the original commands from Git, e.g. '%[2]s push origin main' and 'git push origin main' can be used interchangeably in your project.
	`, Version, Command)
}
