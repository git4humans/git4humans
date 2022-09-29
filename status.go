package git4humans

import (
	"fmt"
	"os"
	"strings"
)

func Status() {
	args := os.Args[2:]

	status := GitStr("status", args...)
	status = strings.ReplaceAll(status, "git push", Command+" publish")
	status = strings.ReplaceAll(status, "git add", Command+" +")
	status = strings.ReplaceAll(status, "git restore", Command+" restore")
	status = strings.ReplaceAll(status, "git commit -a", Command+" save")

	fmt.Print(status)
}
