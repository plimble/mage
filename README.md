# Mage helper

## Installation

```sh
go get -u github.com/plimble/mage/...
$GOPATH/src/github.com/plimble/mage/install
```

## Example

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

func Version() {
	sh.Exec("go version")
}
```

Run

```sh
mage build
mage version
```
