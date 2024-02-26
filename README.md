# Argo [![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/NotWithering/argo/blob/master/LICENSE)

**Argo** is a simple package designed to parse strings into command-line arguments following the POSIX Shell Command Language.

## Example
```go
// cli application
package argo

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NotWithering/argo"
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

		args, incomplete := argo.Parse(in)
		if incomplete {
			fmt.Println("error: incomplete command")
			continue
		}

		fmt.Println(args)
	}
}

```

## What does Argo mean?
**Ar*****g***ument ***G*****o**