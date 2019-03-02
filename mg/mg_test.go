package mg

import (
	"testing"
)

func TestExec(t *testing.T) {
	Exec(`GOOS=linux go build -ldflags="-d -s -w" -a -tags netgo -p="aa bb cc" -installsuffix netgo -o bin/app .`)
}
