package sh

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Run just execute the script.
func Run(script string) error {
	return exec.Command("sh", "-c", script).Run()
}

// Output execute the script, and receive the output as []byte.
func Output(script string, a ...interface{}) ([]byte, error) {
	script = fmt.Sprintf(script, a...)
	return exec.Command("sh", "-c", script).Output()
}

// MustOutput execute the script, receive the output as []byte, panic if error happens.
func MustOutput(script string, a ...interface{}) []byte {
	ret, err := Output(script, a...)
	if err != nil {
		panic(err)
	}
	return ret
}

// String execute the script, and receive the output as string.
func String(script string, a ...interface{}) (result string, err error) {
	bs, err := Output(script, a...)
	if err != nil {
		return
	}
	result = string(bs)
	return
}

// MustString execute the script, receive the output as string. panic if error happens.
func MustString(script string, a ...interface{}) string {
	str, err := String(script, a...)
	if err != nil {
		panic(err)
	}
	return str
}

// Int execute the script, and receive the output as int.
func Int(script string, a ...interface{}) (result int, err error) {
	s, err := String(script, a...)
	if err != nil {
		return
	}
	s = strings.TrimSpace(s)
	result, err = strconv.Atoi(s)
	return
}

// MustInt execute the script, receive the output as int, panic if error happens.
func MustInt(script string, a ...interface{}) int {
	i, err := Int(script, a...)
	if err != nil {
		panic(err)
	}
	return i
}
