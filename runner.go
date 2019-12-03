package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func RunCmd(dir, cmdName string) error {
	fields := strings.Fields(cmdName)

	cmd := exec.Command(fields[0], fields[1:]...)
	cmd.Dir = dir
	cmdReader, err := cmd.StdoutPipe()
	cmdErr, _ := cmd.StderrPipe()
	if err != nil {
		return err
	}
	merged := io.MultiReader(cmdReader, cmdErr)
	scanner := bufio.NewScanner(merged)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Printf("%s\n", text)
		}

	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
