package git4humans

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	GitCommands = []string{"add", "am", "archive", "bisect", "bundle", "checkout", "cherry-pick", "citool", "clean", "clone", "commit", "describe", "diff", "fetch", "format-patch", "gc", "grep", "gui", "init", "log", "maintenance", "merge", "mv", "notes", "pull", "range-diff", "rebase", "reset", "restore", "revert", "rm", "shortlog", "show", "sparse-checkout", "stash", "status", "subdomule", "tag", "worktree", "gitx", "config", "fast-export", "fast-import", "filter-branch", "mergetool", "pack-refs", "prune", "push", "reflog", "remote", "repack", "replace", "annotate", "blame", "bugreport", "count-objects", "difftool", "fsck", "instaweb", "merge-tree", "rerere", "show-branch", "verify-commit", "verify-tag", "whatchanged", "archimport", "cvsexportcommit", "cvsimport", "cvsserver", "imap-send", "p4", "quiltimport", "request-pull", "svn", "revert", "restore", "reset", "update-index", "read-tree", "apply", "checkout-index", "commit-graph", "commit-tree", "hash-object", "index-pack", "merge-file", "merge-index", "mktag", "mktree", "multi-pack-index", "pack-objects", "prune-packed", "symbolic-ref", "unpack-objects", "update-ref", "write-tree", "cat-file", "cherry", "diff-files", "diff-index", "diff-tree", "for-each-ref", "for-each-repo", "tar-commit-id", "ls-files", "ls-remote", "ls-tree", "merge-base", "nave-rev", "pack redundant", "rev-list", "show-index", "show-ref", "unpack-title", "var", "verify-pack", "daemon", "fetch-pack", "http-backend", "send-pack", "update-server-info", "http-fetch", "http-push", "receive-pack", "shell", "upload-archive", "git upload-lack", "check-attr", "check-ignore", "check-mailmap", "check-ref-format", "column", "credential", "credential-cache", "credential-store", "fmt-merge-msg", "hook", "interpre-trailers", "mailinfo", "mainsplit", "merge-one-file", "patch-id", "sh-i18n", "sh-setup", "stripspace"}
)

func Git(command string, args ...string) {
	response := GitStr(command, args...)
	notGit := strings.Contains(response, "fatal: not a git repository")

	if notGit {
		fmt.Printf(`
Err: not a Git repository

Run the following command to start a Git repository:

    %[1]s start 

It will init a fresh Git repository in this directory, then add all files into staging area and do an initial commit.

Use '%[1]s init' if you only want to init a Git repository.
Use '%[1]s + .' if you only want to init and add files into staging.
Use '%[1]s save' if you want to init, add all files, and commit with a specific message.`, Command)

		fmt.Println()
		fmt.Println()
		fmt.Printf("Run '%[1]s start' now? (y/n) ", Command)

		reader := bufio.NewReader(os.Stdin)

		confirm := ""
		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			confirm = strings.Replace(string(input), "\n", "", -1)
		}

		yes := confirm == "Y" || confirm == "y"

		if yes {
			fmt.Println()

			Git("init")
			Git("add", ".")
			Git("commit", "-m", "Initial commit")
		}
	} else {
		isStatus := command == "status" || command == "s"

		if isStatus {
			response = RefineStatus(response)
		}

		fmt.Print(response)
	}
}

func GitStr(command string, args ...string) string {
	params := append([]string{command}, args...)
	cmd := exec.Command("git", params...)

	var out bytes.Buffer
	var err bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &err

	if cmd.Run() != nil {
		return err.String()
	} else {
		return out.String()
	}
}

func NotGit() bool {
	response := GitStr("status")
	notGit := strings.Contains(response, "not a git repository")

	return notGit
}

func HasUpdate() bool {
	response := GitStr("status")

	untracked := strings.Contains(response, "Untracked files:")
	unstaged := strings.Contains(response, "Changes not staged for commit:")

	return untracked || unstaged
}

func HasCommit() bool {
	response := GitStr("status")
	hasCommit := strings.Contains(response, "Changes to be committed")

	return hasCommit
}

func NoCommit() bool {
	return !HasCommit()
}

func RefineStatus(status string) string {
	status = strings.ReplaceAll(status, `use "git push" to publish`, fmt.Sprintf(`use "%[1]s publish" or "%[1]s pub" to publish`, Command))
	status = strings.ReplaceAll(status, "git push", Command+" publish")
	status = strings.ReplaceAll(status, "git add", Command+" +")
	status = strings.ReplaceAll(status, "git restore", Command+" restore")
	status = strings.ReplaceAll(status, "git commit -a", Command+" save")

	return status
}
