package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func sh(command string, args ...interface{}) (string, *exec.ExitError, error) {
	if len(args) > 0 {
		command = fmt.Sprintf(command, args...)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	sOut := stdout.String()
	sErr := stderr.String()

	out := sOut
	if sErr != "" {
		out = sErr
	}

	if err == nil {
		return out, nil, nil
	}

	if e, ok := err.(*exec.ExitError); ok {
		return out, e, nil
	}

	return out, nil, err
}
