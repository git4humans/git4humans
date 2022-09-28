package git4humans

import (
	"bytes"
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
fatal: This project is not a Git repository. 


Use the following command to start a Git repository:

    %[1]s new 


It will init the project as a Git repository, then automatically add all files into staging and do an initial commit.


Use '%[1]s init' if you only want to init a Git repository.

Use '%[1]s +' if you only want to init and add files into staging.

Use '%[1]s save' if you want to init, add all files, and commit with a specific message.
`, Command)
	} else {
		fmt.Println(response)
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
