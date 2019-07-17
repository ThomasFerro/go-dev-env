package docker

import (
	"bufio"
	"log"
	"os/exec"

	"github.com/go-dev-env/builders"
)

// Builder A Docker builder
type Builder struct{}

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
func (d Builder) execute(commandName string, commandArgs ...string) (builders.ArtifactPath, error) {
	args := append([]string{commandName}, commandArgs...)
	cmd := exec.Command("docker", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Docker command error: %v", err)
	}
	log.Printf("Docker command output: %v", string(output))
	return "", err
}

// Build Build the project artifact as a Docker image
func (d Builder) Build(contextPath string) (builders.ArtifactPath, error) {
	return d.execute("build", "-t", "toto", contextPath)
}

// NewBuilder Create a new Docker builder
func NewBuilder() builders.Builder {
	return &Builder{}
}
