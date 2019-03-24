package mg

import (
	"testing"
)

func TestExec(t *testing.T) {
	// ExecX(`zamus migrate-up root@tcp(127.0.0.1:3306)/account?parseTime=true migration`).Run()
	// ExecX(`go build -ldflags="-d -s -w" -a -tags=netgo -installsuffix=netgo -o=bin/app .`).Run()
	BuildLinux(".1", "./bin/app")

}
