package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Everything you wanted to know about codes:
// http://www.umich.edu/~archive/apple2/misc/programmers/vt100.codes.txt

const (
	bold      = 1
	underline = 4
	red       = 31
	green     = 32
	yellow    = 33
	purple    = 34
	fuschia   = 35
	blue      = 36
	white     = 37
)

func decorate(c int, v string) string {
	return fmt.Sprintf("\033[%vm%v\033[0m", c, v)
}

func main() {
	var packageName string
	flag.StringVar(&packageName, "package", "", "the package you want to match (optional)")
	flag.Parse()

	goRoot := strings.ToLower(runtime.GOROOT())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		test := strings.TrimSpace(strings.ToLower(line))
		if !strings.HasPrefix(test, goRoot) {
			if packageName == "" || strings.Contains(line, packageName) {
				line = decorate(bold, line)
			}
		}
		fmt.Println(line) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
