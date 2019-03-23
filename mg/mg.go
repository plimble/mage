package mg

import (
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mg"
)

type Namespace = mg.Namespace

func BuildLinux(path, output string) {
	Exec("go", "build", "-ldflags", "-w -s", "-o", output, path).
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
	env  []string
	dir  string
}

func Exec(cmd string, args ...string) *execBuilder {
	return &execBuilder{
		cmd:  cmd,
		args: args,
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
	}
}

func (b *execBuilder) Env(key, value string) *execBuilder {
	b.env = append(b.env, key+"="+value)
	return b
}

func (b *execBuilder) Dir(path string) *execBuilder {
	b.dir = path
	return b
}

func (b *execBuilder) Run() {
	cmd := exec.Command(b.cmd, b.args...)
	if len(b.env) > 0 {
		cmd.Env = b.env
	}
	if b.dir != "" {
		cmd.Dir = b.dir
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
