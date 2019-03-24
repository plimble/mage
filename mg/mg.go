package mg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mg"
)

type Namespace = mg.Namespace

func BuildLinux(path, output string) {
	Exec("go", "build", "-ldflags=-w -s", "-o", output, path).
		Env("GOOS", "linux").
		Env("GOARCH", "amd64").
		Run()
}

func Build(path, output string) {
	Exec("go", "build", "-o", output, path).Run()
}

func GoGernerate() {
	Exec("go", "generate", "./...").Run()
}

type execBuilder struct {
	cmd  string
	args []string
	env  map[string]string
	dir  string
}

func Exec(cmd string, args ...string) *execBuilder {
	return &execBuilder{
		cmd:  cmd,
		args: args,
		env:  make(map[string]string),
	}
}

func ExecX(cmd string, args ...string) *execBuilder {
	var startQuote bool
	cmdsplits := strings.FieldsFunc(cmd, func(s rune) bool {
		if !startQuote && s == '"' {
			startQuote = true
			return false
		}

		if startQuote && s == '"' {
			startQuote = false
			return false
		}

		if !startQuote && s == ' ' {
			return true
		}

		return false
	})

	for i := 0; i < len(cmdsplits); i++ {
		cmdsplits[i] = strings.ReplaceAll(cmdsplits[i], `"`, "")
	}

	if len(cmdsplits) == 1 {
		return &execBuilder{
			cmd: cmdsplits[0],
		}
	}

	return &execBuilder{
		cmd:  cmdsplits[0],
		args: cmdsplits[1:],
		env:  make(map[string]string),
	}
}

func (b *execBuilder) Env(key, value string) *execBuilder {
	b.env[key] = value
	return b
}

func (b *execBuilder) Dir(path string) *execBuilder {
	b.dir = path
	return b
}

func (b *execBuilder) Run() {
	b.cmd = os.Expand(b.cmd, b.expand)
	for i := range b.args {
		b.args[i] = os.Expand(b.args[i], b.expand)
	}

	cmd := exec.Command(b.cmd, b.args...)
	cmd.Env = os.Environ()
	if len(b.env) > 0 {
		for k, v := range b.env {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
	}
	if b.dir != "" {
		cmd.Dir = b.dir
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		code := ExitStatus(err)
		fmt.Printf("running \"%s %s\" failed with exit code %d\n", b.cmd, strings.Join(b.args, " "), code)
		os.Exit(code)
	}
}

func (b *execBuilder) expand(s string) string {
	s2, ok := b.env[s]
	if ok {
		return s2
	}
	return os.Getenv(s)
}

type exitStatus interface {
	ExitStatus() int
}

func ExitStatus(err error) int {
	if err == nil {
		return 0
	}
	if e, ok := err.(exitStatus); ok {
		return e.ExitStatus()
	}
	if e, ok := err.(*exec.ExitError); ok {
		if ex, ok := e.Sys().(exitStatus); ok {
			return ex.ExitStatus()
		}
	}
	return 1
}
