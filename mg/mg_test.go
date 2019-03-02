package mg

import (
	"testing"
)

func TestExec(t *testing.T) {
	Exec(`go build -ldflags="-d -s -w" -a -tags=netgo -installsuffix=netgo -o=bin/app .`)
}
