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
    
    %[1]s %[2]s file
    %[1]s %[2]s file1 file2 file3
    %[1]s %[2]s dir/file
    %[1]s %[2]s file1 dir/file2
            `, Command, command)
		}
	}
}
