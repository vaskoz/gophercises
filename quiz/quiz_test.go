package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"testing"
)

func TestMainBadArgs(t *testing.T) {
	args = []string{"quiz", "-limit", "-1"}
	buff := new(bytes.Buffer)
	stderr = buff
	var code int
	exit = func(c int) {
		code = c
	}
	main()
	if code != 1 {
		t.Errorf("expected exit code 1, but got %d", code)
	}
	if out := buff.String(); !strings.Contains(out, "invalid value") {
		t.Errorf("expected an invalid value message, but got %s", out)
	}
}

func TestMainMissingCSVFile(t *testing.T) {
	args = []string{"quiz", "-limit", "3", "-csv", "legitfile"}
	buff := new(bytes.Buffer)
	stderr = buff
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
	if out := buff.String(); !strings.Contains(out, "Could not open csv file") {
		t.Errorf("expected file could not be opened, but got %s", out)
	}
}

func TestMainBadCSVReadAll(t *testing.T) {
	args = []string{"quiz", "-limit", "3", "-csv", "legitfile"}
	buff := new(bytes.Buffer)
	stderr = buff
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
	if out := buff.String(); !strings.Contains(out, "error in reading data file") {
		t.Errorf("expected error reading file, but got %s", out)
	}
}

func TestMainImmediateTimeout(t *testing.T) {
	buff := new(bytes.Buffer)
	stdout = buff
	stdin = strings.NewReader("10 2 11 3 14 4 5 6 5 6 6 7")
	args = []string{"quiz", "-limit", "0", "-csv", "legitfile"}
	var code int
	exit = func(c int) {
		code = c
	}
	open = func(f string) (io.Reader, error) {
		return strings.NewReader(`5+5,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7`), nil
	}
	main()
	if code != 0 {
		t.Errorf("Expected a successful run, but got %d", code)
	}
	out := buff.String()
	if !strings.Contains(out, "You scored 0 out of 12.") {
		t.Errorf("Expected immediate timeout message, but got %s", out)
	}
}

func TestMainFullMarksUnit(t *testing.T) {
	buffOutput := new(bytes.Buffer)
	stdout = buffOutput
	stdin = strings.NewReader("10\n2\n11\n3\n14\n4\n5\n6\n5\n6\n6\n7\n")
	args = []string{"quiz"}
	var code int
	exit = func(c int) {
		code = c
	}
	open = func(f string) (io.Reader, error) {
		return strings.NewReader(`5+5,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7`), nil
	}
	main()
	if code != 0 {
		t.Errorf("Expected a successful run, but got %d", code)
	}
	out := buffOutput.String()
	if !strings.Contains(out, "You scored 12 out of 12.") {
		t.Errorf("Expected successful run, but got %s", out)
	}
}

func TestMainBonusCaseInsensitive(t *testing.T) {
	buffOutput := new(bytes.Buffer)
	stdout = buffOutput
	stdin = strings.NewReader("bar\nho\n")
	args = []string{"quiz"}
	var code int
	exit = func(c int) {
		code = c
	}
	open = func(f string) (io.Reader, error) {
		return strings.NewReader(`foo,bAr
hey,HO`), nil
	}
	main()
	if code != 0 {
		t.Errorf("Expected a successful run, but got %d", code)
	}
	out := buffOutput.String()
	if !strings.Contains(out, "You scored 2 out of 2.") {
		t.Errorf("Expected successful run, but got %s", out)
	}
}

func TestMainBonusSpaceInsensitive(t *testing.T) {
	buffOutput := new(bytes.Buffer)
	stdout = buffOutput
	stdin = strings.NewReader(" bar \n ho \n")
	args = []string{"quiz"}
	var code int
	exit = func(c int) {
		code = c
	}
	open = func(f string) (io.Reader, error) {
		return strings.NewReader(`foo,bAr
hey,HO`), nil
	}
	main()
	if code != 0 {
		t.Errorf("Expected a successful run, but got %d", code)
	}
	out := buffOutput.String()
	if !strings.Contains(out, "You scored 2 out of 2.") {
		t.Errorf("Expected successful run, but got %s", out)
	}
}

func TestMainBonusShuffle(t *testing.T) {
	rand.Seed(1)
	buffOutput := new(bytes.Buffer)
	stdout = buffOutput
	stdin = strings.NewReader("1\n2\n3\n4\n5\n")
	args = []string{"quiz", "-shuffle"}
	var code int
	exit = func(c int) {
		code = c
	}
	open = func(f string) (io.Reader, error) {
		return strings.NewReader(`a,1
b,2
c,3
d,4
e,5`), nil
	}
	main()
	if code != 0 {
		t.Errorf("Expected a successful run, but got %d", code)
	}
	out := buffOutput.String()
	if !strings.Contains(out, "Problem #1: b = Problem #2: a = Problem #3: d = Problem #4: c = Problem #5: e =") {
		t.Errorf("Expected shuffle to prevent perfect score, but got %s", out)
	}
}
