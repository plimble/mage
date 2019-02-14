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

type Build mg.Namespace

func (Build) Linux() {
	sh.BuildLinux(".", "./bin/app-linux")
	fmt.Println("Build Done")
}

func (Build) Mac() {
	sh.BuildMac(".", "./bin/app-mac")
	fmt.Println("Build Done")
}

func (Build) All() {
	Build{}.Linux()
	Build{}.Mac()
	fmt.Println("Build Done")
}

func Version() {
	sh.Exec("go version")
}
```

Run

```sh
mage build:linux
mage build:mac
mage build:all
mage version
```
