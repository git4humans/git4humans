package main

import (
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
		gh.New()
	case "+":
		gh.Add()
	case "-":
		gh.Delete()
	case "del":
		gh.Delete()
	case "delete":
		gh.Delete()
	case "rename":
		gh.Move()
	case "move":
		gh.Move()
	case "copy":
		gh.Copy()
	case "save":
		gh.Save()
	case "push":
		gh.Push()
	case "sync":
		gh.Sync()
	case "branch":
		gh.Branch()
	case "switch":
		gh.SwitchBranch()
	case "repo":
		gh.Repo()
	case "pr":
		gh.Pr()
	case "user":
		gh.User()
	case "help":
		gh.Help()
	default:
		if gh.Contains(gitCommands, subcommand) {
			git(subcommand, args...)
		} else {
			fmt.Printf(`Error: '%[1]s' is not a valid command. See %[2]s help.
            `, subcommand, command)
		}
	}
}
