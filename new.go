package git4humans

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func New() {
	args := os.Args[2:]
	dir := ""

	if len(args) > 0 {
		dir = args[0]
	}

	if len(dir) <= 0 {
		fmt.Print("Working directory: ")

		reader := bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadLine()

		if len(input) > 0 {
			dir = strings.Replace(string(input), "\n", "", -1)
		}
	}

	if len(dir) <= 0 {
		fmt.Println("Error: you should specify a working directory for the new Git repository.")
		return
	}

	err := os.Mkdir(dir, os.ModePerm)

	if err != nil {
		fmt.Println("Error: failed creating directory for the new Git repository.")
		return
	}

	err = os.Chdir(dir)

	if err != nil {
		panic(err)
	} else {
		Git("init")
	}
}
