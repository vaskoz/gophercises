package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestIntegration(t *testing.T) {
	setup()
	args = []string{"quiz"}
	stdin = strings.NewReader("10 2 11 3 14 4 5 6 5 6 6 7")
	exit = func(c int) {}
	buff := new(bytes.Buffer)
	stdout = buff
	main()
	if out := buff.String(); !strings.Contains(out, "You scored 12 out of 12") {
		t.Errorf("Integration test failed")
	}
}
