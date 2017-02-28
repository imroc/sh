package sh

import (
	"os/exec"
	"strconv"
	"strings"
)

// Run just execute the script.
func Run(script string) error {
	return exec.Command("sh", "-c", script).Run()
}

// Output execute the script, and receive the output as []byte.
func Output(script string) ([]byte, error) {
	return exec.Command("sh", "-c", script).Output()
}

// String execute the script, and receive the output as string.
func String(script string) (result string, err error) {
	bs, err := Output(script)
	if err != nil {
		return
	}
	result = string(bs)
	return
}

// Int execute the script, and receive the output as int.
func Int(script string) (result int, err error) {
	s, err := String(script)
	if err != nil {
		return
	}
	s = strings.TrimSpace(s)
	result, err = strconv.Atoi(s)
	return
}
