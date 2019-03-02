package mg

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Namespace = mg.Namespace

type opts struct {
	env map[string]string
}

type Options func(o *opts)

func WithEnv(key, value string) Options {
	return func(o *opts) {
		if o.env == nil {
			o.env = make(map[string]string)
		}

		o.env[key] = value
	}
}

func Exec(cmd string, options ...Options) {
	o := &opts{}
	for _, opt := range options {
		opt(o)
	}

	var err error
	cmdsplits := strings.Split(cmd, " ")
	for i := 0; i < len(cmdsplits); i++ {
		cmdsplits[i] = strings.ReplaceAll(cmdsplits[i], `"`, "")
	}
	if len(cmdsplits) == 1 {
		_, err = sh.Exec(o.env, os.Stdout, os.Stderr, cmdsplits[0])
	} else {
		_, err = sh.Exec(o.env, os.Stdout, os.Stderr, cmdsplits[0], cmdsplits[1:]...)
	}

	exitn := sh.ExitStatus(err)

	if exitn > 0 {
		os.Exit(exitn)
	}
}

func BuildLinux(path, output string) {
	Exec(fmt.Sprintf("go build -a -o=%s %s", output, path),
		WithEnv("GOOS", "linux"),
	)
}

func Build(path, output string) {
	Exec(fmt.Sprintf("go build -a -o=%s %s", output, path))
}

func GoGernerate() {
	Exec("go generate ./...")
}
