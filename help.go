package git4humans

import "fmt"

func Help() {
	fmt.Printf(`
Git for Humans version %[1]s

usage: %[2]s [<options>] <command> [<args]

These are common commands you will most likely use in everyday situation.

Start a working area
    copy  	Copy an existing repository to a new directory
    new		Create an Git repository in a new directory 
    start	Initialize an fresh Git repository in the current working directory

Work on the current change 
    + 		Add file contents to the index 
    - 		Delete file(s) from the directory and from the index 
    move 	Move a file, a directory, or a symlink 
    rename 	Rename a file, a directory, or a symlink 
    restore	Restore deleted file(s)
    save 	Record changes to the repository

Collaborate with others
    repo 	List, add, or delete remote repositories
    sync 	Fetch from and integrate with another repository or a local branch 
    submit 	Update changes into remote repository 
    fetch 	Download from another repository

Grow, mark, and tweak your common history 
    branch 	List, create, or delete branches 
    switch 	Switch to another branch 
    merge 	Join two or more branches 
    reset 	Reset current HEAD to a specified state
    tag 	Create, list, delete, or verify a tag object

Examine history and state 
    status 	Show the working tree status (shortcut: '%[2]s s') 
    log 	Show commit logs (shortcut: '%[2]s l')
    diff 	Show changes between commits, commit and working tree, etc.

Configurations 
    config	Show or update Git configuration
    user	Show or update Git user configuration (local or global)

Also, you can use other standard commands that comes with Git. 
As an example, '%[2]s push origin main' will work exactly as 'git push origin main' 
	`, Version, Command)
}
