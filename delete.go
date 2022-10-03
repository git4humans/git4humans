package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Remove file(s)
func Delete() {
	if IsHelp() {
		DeleteHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	if len(args) > 0 {
		Git("rm", args...)
		Git("status")
	} else {
		files := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Println()
		fmt.Print("File(s) to delete: ")

		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			files = strings.Replace(string(input), "\n", "", -1)
		}

		if len(files) > 0 {
			Git("rm", strings.Fields(files)...)
			Git("status")
		} else {
			fmt.Printf(`
You should specify file(s) to delete.

Examples:

    %[1]s - file.txt
    %[1]s - file1.txt file2.txt file3.txt
    %[1]s - file1.txt dir2/file2.txt
            `, Command, command)
		}
	}
}

func DeleteHelp() {
	command := os.Args[1]

	fmt.Printf(`
Delete file(s) at the given path(s), then stages the deletion for the next commit.

usage: %[1]s %[2]s [<file>]

examples:

    %[1]s %[2]s file.txt
    %[1]s %[2]s file1.txt file2.txt file3.txt
    %[1]s %[2]s file1.txt dir2/file2.txt
	`, Command, command)
}
