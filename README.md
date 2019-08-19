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

	"github.com/plimble/mage/mg"
)

type Build mg.Namespace

func (Build) Linux() {
	// build with cgo
	mg.BuildLinux(".", "./bin/app-linux", true)
	// build with disabled cgo
	mg.BuildLinux(".", "./bin/app-linux", false)
	fmt.Println("Build Done")
}

func (Build) Mac() {
	mg.BuildMac(".", "./bin/app-mac")
	fmt.Println("Build Done")
}

func (Build) All() {
	Build{}.Linux()
	Build{}.Mac()
	fmt.Println("Build Done")
}

func Version() {
	mg.Exec("go version")
}
```

Run

```sh
mage build:linux
mage build:mac
mage build:all
mage version
```
