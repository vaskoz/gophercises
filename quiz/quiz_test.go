package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestMainBadArgs(t *testing.T) {
	args = []string{"quiz", "-limit", "-1"}
	var code int
	exit = func(c int) {
		code = c
	}
	main()
	if code != 1 {
		t.Errorf("expected exit code 1, but got %d", code)
	}
}

func TestMainMissingCSVFile(t *testing.T) {
	args = []string{"quiz", "-limit", "3", "-csv", "legitfile"}
	var code int
	exit = func(c int) {
		code = c
	}
	open = func(f string) (io.Reader, error) {
		return nil, fmt.Errorf("can't open file")
	}
	main()
	if code != 1 {
		t.Errorf("missing file should return an error code of 1, but got %d", code)
	}
}

func TestMainBadCSVReadAll(t *testing.T) {
	args = []string{"quiz", "-limit", "3", "-csv", "legitfile"}
	var code int
	exit = func(c int) {
		code = c
	}
	open = func(f string) (io.Reader, error) {
		return strings.NewReader(`"Rob","Pike",rob,"blah"
Ken,Thompson,ken`), nil
	}
	main()
	if code != 1 {
		t.Errorf("incorrect CSV should return an error code of 1, but got %d", code)
	}
}
