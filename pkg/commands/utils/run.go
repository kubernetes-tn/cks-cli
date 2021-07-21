package utils

import (
	"fmt"
	"os/exec"
)

func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	// error output and standard output of command is connected to the same pipeline
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// Get real-time output from the pipe to the terminal and print
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
