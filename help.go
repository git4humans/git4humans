package git4humans

import "fmt"

func Help() {
	fmt.Printf(`
Git for Humans version %[1]s

usage: %[2]s [<options>] <command> [<args]

These are common commands you will most likely use in everyday situation.

Starting point
    copy  	Copy an existing Git repository to a new directory (short: '%[2]s cp')
    new		Create a new Git repository in a new directory (short: '%[2]s n')
    start	Start a new Git repository in the current working directory (short: '%[2]s s')
    refresh     Refresh an existing Git repository (short: '%[2]s r')

Managing changes 
    + 		Add file(s) to the staging area for inclusion in the next commit
    - 		Delete file(s), then stages the deletion for the next commit
    move 	Move a file, a directory, or a symlink (short: '%[2]s mv')
    rename 	Rename a file, a directory, or a symlink (short: '%[2]s ren')
    restore	Restore deleted file(s) (short: '%[2]s rs')
    save 	Record changes in your local repository (short: '%[2]s sv')

Collaborations
    repo 	Manage remote repositories (list, add, delete, etc.)
    sync 	Download from and integrate with a repository or a local branch 
    up          Upload your local changes to a remote repository

Branches and history 
    branch 	Manage branches (list, create, delete, etc.) 
    switch 	Switch from your current branch to another branch 
    merge 	Join two or more branches
    reset 	Reset current HEAD to a specified state
    tag 	Create, list, delete, or verify a tag object

Examine history and state 
    st 	    Show the status of the current working project
    sst     Show the short status of the current working project 
    lg 	    Show a reversed-ordered list of commits (logs)
    slg     Show a list of commits in a simple oneline format (simple logs)
    last    Show the last commit in the log history

Configurations 
    config	Show or update configuration
    user	Show or update user configuration (local or global)

You may also run all the original commands from Git. For example, '%[2]s push origin main' and 'git push origin main' can be used interchangeably in your project.
	`, Version, Command)
}
