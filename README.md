# Argo
[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/notwithering/argo/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/notwithering/argo)](https://goreportcard.com/report/github.com/notwithering/argo)

**Argo** is a simple package designed to parse strings into command-line arguments following the POSIX Shell Command Language.

## Example
```go
// cli application
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/notwithering/argo"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(" > ")

		in, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		in = strings.TrimSpace(in)

		args, incomplete := argo.Parse(in)
		if incomplete {
			fmt.Println("error: incomplete command")
			continue
		}

		for i, a := range args {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Printf("\"%s\"", a)
		}
		fmt.Print("\n")

		if len(args) > 0 {
			if args[0] == "exit" {
				return
			}
		}
	}
}
```

## What does Argo mean?
**Ar*****g***ument ***G*****o**