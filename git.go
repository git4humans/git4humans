package git4humans

import (
	"fmt"
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
Err: this project is not a Git repository

Use the following command to start a Git repository:

    %[1]s start 

That will create a fresh Git repository in your project, then automatically stage all files and do initial commit.

Use '%[1]s init' if you only want to start a fresh Git repository.
Use '%[1]s + .' if you only want to start Git and stage all files in your project.
Use '%[1]s save' if you want to start Git, stage all files, and commit with a specific message.`, Command)

		fmt.Println()
	} else {
		if command == "status" || command == "s" {
			response = RefineStatus(response)
		}

		fmt.Print(response)
	}
}

func GitStr(command string, args ...string) string {
	params := append([]string{command}, args...)
	cmd := exec.Command("git", params...)

	out, _ := cmd.CombinedOutput()
	result := string(out)

	return result
}

func NotGit() bool {
	response := GitStr("branch", "--show-current")
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
	status = strings.ReplaceAll(status, "git rm", Command+" -")
	status = strings.ReplaceAll(status, "git restore", Command+" restore")
	status = strings.ReplaceAll(status, "git commit -a", Command+" save")

	return status
}

func ListChanges() {
	status := GitStr("status")

	lines := strings.Split(strings.ReplaceAll(status, "\r\n", "\n"), "\n")
	result := []string{}

	for _, line := range lines {
		//onBranch := strings.HasPrefix(line, "On branch")
		isUptodate := strings.HasPrefix(line, "Your branch is up to date")
		isAhead := strings.HasPrefix(line, "Your branch is ahead of")
		noCommit := strings.HasPrefix(line, "No commits yet")
		noChanges := strings.HasPrefix(line, "no changes added to commit")
		noAdded := strings.HasPrefix(line, "nothing added to commit")
		useAdd := strings.Contains(line, `use "git add <file>`)
		useRestore := strings.Contains(line, `use "git restore <file>`)
		useRemove := strings.Contains(line, `use "git rm`)
		usePush := strings.Contains(line, `use "git push"`)

		hide := (noCommit || noChanges || noAdded || useAdd || useRestore || useRemove || usePush || isUptodate || isAhead)
		show := !hide

		if show {
			result = append(result, line)
		}
	}

	fmt.Print(strings.Join(result, "\n"))
}
