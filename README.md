# Mage helper

```go
// +build mage

package main

import (
	"fmt"

	"github.com/plimble/mage/sh"
)

func Build() {
	sh.BuildLinux(".", "./bin/app")
	fmt.Println("Build Done")
}

func Deploy() {
	Build()
	sh.Exec("serverless deploy -v")
}

func Remove() {
	sh.Exec("serverless remove -v")
}
```
