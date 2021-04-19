package main

import (
	"log"
	"os/exec"
)

func listAllProcess() string {
	cmd := exec.Command("tasklist")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
