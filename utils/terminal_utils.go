package utils

import (
	"runtime"
	"os/exec"
	"bytes"
)

func ExecuteCommand(input string) (out string, err error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(getShell(), getCommandArg(), input)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	out = stdout.String() + stderr.String()
	return out, err
}

func getShell() string {
	if runtime.GOOS == "windows" {
		return "cmd"
	}
	return "sh"
}

func getCommandArg() string {
	if runtime.GOOS == "windows"{
		return "/C"
	} 
	return "-c"
}