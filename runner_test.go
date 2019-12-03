package cmd

import (
	"log"
	"testing"
)

func TestRunCmd(t *testing.T) {
	err := RunCmd("/", "echo 'ok'")
	if err != nil {
		log.Fatal(err)
	}
}
