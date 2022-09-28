package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Remove file(s)
func Delete() {
	command := os.Args[1]
	args := os.Args[2:]

	if len(args) > 0 {
		Git("rm", args...)
		Git("status")
	} else {
		files := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("\nFile(s) to delete: ")

		text, _, _ := reader.ReadLine()

		if len(text) > 0 {
			files = strings.Replace(string(text), "\n", "", -1)
		}

		if len(files) > 0 {
			Git("rm", files)
			Git("status")
		} else {
			fmt.Printf(`
Error: You should specify file(s) to delete.

Examples:
    
    %[1]s %[2]s file
    %[1]s %[2]s file1 file2 file3
    %[1]s %[2]s dir/file
    %[1]s %[2]s file1 dir/file2
            `, Command, command)
		}
	}
}
