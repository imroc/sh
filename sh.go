package sh

import "os/exec"

func Run(script string) error {
	return exec.Command("sh", "-c", script).Run()
}

func Output(script string) ([]byte, error) {
	return exec.Command("sh", "-c", script).Output()
}

func String(script string) (result string, err error) {
	bs, err := Output(script)
	if err != nil {
		return
	}
	result = string(bs)
	return
}
