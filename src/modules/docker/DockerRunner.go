package docker

import (
	"bufio"
	"log"
	"os/exec"
)

func getStdoutScanner(cmd *exec.Cmd) (*bufio.Scanner, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Docker exec: Could not connect to the standard output: %v", err)
		return nil, err
	}

	return bufio.NewScanner(stdout), nil
}

func getStderrScanner(cmd *exec.Cmd) (*bufio.Scanner, error) {
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Printf("Docker exec: Could not connect to the standard error output: %v", err)
		return nil, err
	}

	return bufio.NewScanner(stderr), nil
}

// TODO : How to return the results ?
func Execute(commandName string, commandArgs ...string) error {
	// TODO : Get the error pipe too
	// TODO : Docker build
	// TODO : Extract in a Docker module
	args := append([]string{commandName}, commandArgs...)
	cmd := exec.Command("docker", args...)

	outScanner, _ := getStdoutScanner(cmd)
	errScanner, _ := getStderrScanner(cmd)

	if err := cmd.Start(); err != nil {
		log.Printf("Docker command failed: %v", err)
		return err
	}

	for outScanner.Scan() {
		log.Printf("Output from the executed Docker command: %v", outScanner.Text())
	}

	if err := outScanner.Err(); err != nil {
		log.Printf("Docker command failed: %s", err)
		return err
	}

	for errScanner.Scan() {
		log.Printf("Error output from the executed Docker command: %v", errScanner.Text())
	}

	if err := errScanner.Err(); err != nil {
		log.Printf("Docker command failed: %s", err)
		return err
	}

	log.Printf("Docker command execution finished")
	return nil
}
