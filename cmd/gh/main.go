package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	gh "github.com/git4humans/git4humans"
)

var (
	command     = "gh"
	gitCommands = []string{"add", "am", "archive", "bisect", "bundle", "checkout", "cherry-pick", "citool", "clean", "clone", "commit", "describe", "diff", "fetch", "format-patch", "gc", "grep", "gui", "init", "log", "maintenance", "merge", "mv", "notes", "pull", "range-diff", "rebase", "reset", "restore", "revert", "rm", "shortlog", "show", "sparse-checkout", "stash", "status", "subdomule", "tag", "worktree", "gitx", "config", "fast-export", "fast-import", "filter-branch", "mergetool", "pack-refs", "prune", "reflog", "remote", "repack", "replace", "annotate", "blame", "bugreport", "count-objects", "difftool", "fsck", "instaweb", "merge-tree", "rerere", "show-branch", "verify-commit", "verify-tag", "whatchanged", "archimport", "cvsexportcommit", "cvsimport", "cvsserver", "imap-send", "p4", "quiltimport", "request-pull", "svn", "revert", "restore", "reset", "update-index", "read-tree", "apply", "checkout-index", "commit-graph", "commit-tree", "hash-object", "index-pack", "merge-file", "merge-index", "mktag", "mktree", "multi-pack-index", "pack-objects", "prune-packed", "symbolic-ref", "unpack-objects", "update-ref", "write-tree", "cat-file", "cherry", "diff-files", "diff-index", "diff-tree", "for-each-ref", "for-each-repo", "tar-commit-id", "ls-files", "ls-remote", "ls-tree", "merge-base", "nave-rev", "pack redundant", "rev-list", "show-index", "show-ref", "unpack-title", "var", "verify-pack", "daemon", "fetch-pack", "http-backend", "send-pack", "update-server-info", "http-fetch", "http-push", "receive-pack", "shell", "upload-archive", "git upload-lack", "check-attr", "check-ignore", "check-mailmap", "check-ref-format", "column", "credential", "credential-cache", "credential-store", "fmt-merge-msg", "hook", "interpre-trailers", "mailinfo", "mainsplit", "merge-one-file", "patch-id", "sh-i18n", "sh-setup", "stripspace"}
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		execute()
	} else {
		gh.Help()
	}
}

func gitStr(subcommand string, args ...string) string {
	params := append([]string{subcommand}, args...)
	cmd := exec.Command("git", params...)

	var out bytes.Buffer
	var err bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &err

	if cmd.Run() != nil {
		error := err.String()
		notGit := strings.Contains(error, "not a git repository")

		if notGit {
			error = fmt.Sprintf(`
%[1]s
Use the following command to create a new Git repository:

    %[2]s new 


Note:
The '%[2]s new' will initialize a Git repository, then add all files into staging, and do initial commit.

Use '%[2]s init' if you only want to initialize.
Use '%[2]s +' if you only want to initialize and add files into staging.
Use '%[2]s save' if you only want to initialize, add all files, and commit with your own message.`, error, command)
		}

		return error
	} else {
		return out.String()
	}
}

func git(subcommand string, args ...string) {
	fmt.Println(gitStr(subcommand, args...))
}

func notGit() bool {
	status := gitStr("status")
	result := strings.Contains(status, "not a git repository")

	return result
}

func hasUpdate() bool {
	status := gitStr("status")

	untracked := strings.Contains(status, "Untracked files:")
	unstaged := strings.Contains(status, "Changes not staged for commit:")

	return untracked || unstaged
}

func execute() {
	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "new":
		new()
	case "+":
		add()
	case "-":
		del()
	case "del":
		del()
	case "delete":
		del()
	case "rename":
		move()
	case "move":
		move()
	case "copy":
		copy()
	case "save":
		save()
	case "push":
		push()
	case "sync":
		sync()
	case "branch":
		branch()
	case "switch":
		switchBranch()
	case "repo":
		repo()
	case "pr":
		pr()
	case "user":
		user()
	case "help":
		gh.Help()
	default:
		if contains(gitCommands, subcommand) {
			git(subcommand, args...)
		} else {
			fmt.Printf(`Error: '%[1]s' is not a valid command. See %[2]s help.
            `, subcommand, command)
		}
	}
}

// Initializing a new git repository:
//
// git init
// git add .
// git commit
func new() {
	fmt.Print(`
This action will initialize a new Git repository, then add all files into staging and do initial commit:

    git init 
    git add .
    git commit -m "Initial commit"

`)
	fmt.Print("Continue? (y/n) ")

	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()

	if len(text) > 0 {
		confirm := strings.Replace(string(text), "\n", "", -1)
		yes := confirm == "Y" || confirm == "y"

		if yes {
			args := os.Args[2:]

			git("init", args...)
			git("add", ".")
			git("commit", "-m", "Initial commit")
		}
	}
}

// Add untracked files into staging.
// Automatically initialize if the repo is not initialized.
func add() {
	args := os.Args[2:]

	if notGit() {
		git("init")
	}

	if len(args) > 0 {
		git("add", args...)
	} else {
		git("add", ".")
	}

	git("status")
}

// Remove file(s)
func del() {
	subcommand := os.Args[1]
	args := os.Args[2:]

	if len(args) > 0 {
		git("rm", args...)
		git("status")
	} else {
		files := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("\nFile(s) to delete: ")

		text, _, _ := reader.ReadLine()

		if len(text) > 0 {
			files = strings.Replace(string(text), "\n", "", -1)
		}

		if len(files) > 0 {
			git("rm", files)
			git("status")
		} else {
			fmt.Printf(`
Error: You should specify file(s) to delete.

Examples:
    
    %[1]s %[2]s file
    %[1]s %[2]s file1 file2 file3
    %[1]s %[2]s dir/file
    %[1]s %[2]s file1 dir/file2
            `, command, subcommand)
		}
	}
}

// Move or rename a file
func move() {
	command := os.Args[1]
	args := os.Args[2:]

	if len(args) >= 2 {
		oldName := args[0]
		newName := args[1]
		options := args[2:]

		if command == "rename" {
			git("mv", append([]string{oldName, newName}, options...)...)
		} else {
			git("mv", oldName, newName)
		}
	}
}

// Clone a remote repository
func copy() {
	args := os.Args[2:]

	if len(args) > 0 {
		git("clone", args...)
	} else {
		url := ""
		dir := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("\nRepository URL: ")

		textUrl, _, _ := reader.ReadLine()

		if len(textUrl) > 0 {
			url = strings.Replace(string(textUrl), "\n", "", -1)
		}

		fmt.Print("Target directory: ")

		textDir, _, _ := reader.ReadLine()

		if len(textDir) > 0 {
			dir = strings.Replace(string(textDir), "\n", "", -1)
		}

		fmt.Println("")

		if len(url) > 0 {
			if len(dir) > 0 {
				git("clone", url, dir)
			} else {
				git("clone", url)
			}
		}
	}
}

func save() {
	args := os.Args[2:]
	params := remove(args, "-m")
	message := ""

	status := gitStr("status")
	noCommit := strings.Contains(status, "nothing to commit")

	if noCommit {
		fmt.Println(status)
		return
	}

	if len(params) > 0 {
		message = params[0]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nMessage: ")

		text, _, _ := reader.ReadLine()

		if len(text) > 0 {
			message = strings.Replace(string(text), "\n", "", -1)
		}
	}

	if len(message) > 0 {
		if notGit() {
			git("init")
		}

		if hasUpdate() {
			git("add", ".")
		}

		git("commit", "-m", message)
	} else {
		fmt.Println("\nUnable to save a commit without message.")
	}
}

func push() {

}

func sync() {
	args := os.Args[2:]

	git("pull", args...)
}

func branch() {
	args := os.Args[2:]

	if len(args) > 0 {
		branch := args[0]

		fmt.Println("Creating branch " + branch + "...")
		git("branch", branch)

		fmt.Println("Switching to " + branch + "...")
		git("switch", branch)

		showBranch()
	} else {
		showBranch()
	}
}

func showBranch() {
	branch := gitStr("branch", "--show-current")

	fmt.Println("On branch " + branch)
	fmt.Println()
	fmt.Println("List branch:")

	git("branch")
}

func switchBranch() {
	args := os.Args[2:]

	message := gitStr("switch", args...)
	error := strings.Contains(message, "fatal:")

	if error {
		fmt.Println(message)
	} else {
		showBranch()
	}
}

func repo() {
	args := os.Args[2:]

	if len(args) > 0 {
		command := args[0]

		if command == "rename" {
			renameRepo()
		} else if command == "show" {
			showRepo()
		} else if command == "prune" {
			pruneRepo()
		} else if command == "url" {
			urlRepo()
		} else if command == "delete" {
			deleteRepo()
		} else if command == "del" {
			deleteRepo()
		} else if command == "-" {
			deleteRepo()
		} else {
			addRepo()
		}
	} else {
		listRepo()
	}
}

func hasRepo(repo string) bool {
	repos := gitStr("remote", "-v")
	return strings.Contains(repos, repo)
}

func addRepo() {
	command := os.Args[2]
	name := "origin"
	url := ""

	var args []string

	if command == "+" {
		args = os.Args[3:]
	} else {
		args = os.Args[2:]
	}

	if len(args) >= 2 {
		name = args[0]
		url = args[1]
	} else if len(args) >= 1 {
		url = args[0]
	}

	if len(url) > 0 {
		if hasRepo(name) {
			gitStr("remote", "rm", name)
		}

		git("remote", "add", name, url)
	}

	listRepo()
}

func deleteRepo() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	if hasRepo(name) {
		message := gitStr("remote", "rm", name)
		error := strings.Contains(message, "fatal:")

		if error {
			fmt.Println(message)
		} else {
			fmt.Println("Remote repository " + name + " has been removed.")

			git("remote", "-v")
		}
	} else {
		fmt.Println("Remote repository " + name + " is not defined.")
		fmt.Println()

		git("remote", "-v")
	}
}

func renameRepo() {
	args := os.Args[3:]

	oldName := "origin"
	newName := ""

	if len(args) >= 2 {
		oldName = args[0]
		newName = args[1]
	} else if len(args) >= 1 {
		newName = args[0]
	}

	if oldName != "" && newName != "" {
		message := gitStr("remote", "rename", oldName, newName)
		error := strings.Contains(message, "fatal:")

		if error {
			fmt.Println(message)
		} else {
			fmt.Println(message)

			git("remote", "-v")
		}
	}
}

func showRepo() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	git("remote", "show", name)
}

func pruneRepo() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	git("remote", "prune", name)
}

func urlRepo() {
	args := os.Args[3:]
	name := "origin"

	if len(args) > 0 {
		name = args[0]
	}

	git("remote", "get-url", name)
}

func listRepo() {
	message := gitStr("remote", "-v")

	if len(message) > 0 {
		fmt.Println(message)
	} else {
		fmt.Printf(`
Remote repositories are empty.

Use the following command to add repository:
    
    %[1]s repo <url>
    %[1]s repo <name> <url>
    %[1]s repo + <url>
    %[1]s repo + <name> <url>

Examples:

    %[1]s repo https://github.com/pytorch/pytorch
    %[1]s repo origin https://github.com/pytorch/pytorch
    %[1]s repo + https://github.com/pytorch/pytorch
    %[1]s repo + origin https://github.com/pytorch/pytorch
    `, command)
	}
}

func pr() {
	args := os.Args[2:]

	git("request-pull", args...)
}

func user() {
	args := os.Args[2:]
	global := contains(args, "--global")
	params := remove(args, "--global")

	if len(params) >= 2 {
		name := params[0]
		email := params[1]

		if global {
			git("config", "--global", "user.name", name)
			git("config", "--global", "user.email", email)
		} else {
			git("config", "user.name", name)
			git("config", "user.email", email)
		}

		fmt.Println()
	}

	if global {
		name := gitStr("config", "--global", "user.name")
		email := gitStr("config", "--global", "user.email")

		fmt.Printf(`
Git user configuration (global):

user.name   %[1]suser.email  %[2]s
        `, name, email)
	} else {
		name := gitStr("config", "user.name")
		email := gitStr("config", "user.email")

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
