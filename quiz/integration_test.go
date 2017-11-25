package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestIntegration(t *testing.T) {
	setup()
	args = []string{"quiz"}
	stdin = strings.NewReader("10\n2\n11\n3\n14\n4\n5\n6\n5\n6\n6\n7\n")
	exit = func(c int) {}
	buff := new(bytes.Buffer)
	stdout = buff
	main()
	if out := buff.String(); !strings.Contains(out, "You scored 12 out of 12") {
		t.Errorf("Integration test failed")
	}
}
