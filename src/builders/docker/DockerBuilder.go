package docker

import (
	"bufio"
	"log"
	"os/exec"
	"regexp"

	"github.com/go-dev-env/builders"
)

// Builder A Docker builder
type Builder struct {
	artifactName string
}

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

func extractArtifact(output string) builders.ArtifactPath {
	log.Println("Extracting artifact name")
	r := regexp.MustCompile("Successfully tagged (.*)")
	submatches := r.FindStringSubmatch(output)

	if len(submatches) > 1 {
		log.Printf("Artifact name %v successfully extracted", submatches[1])
		return builders.ArtifactPath(submatches[1])
	}

	log.Println("Could not find any artifact name")
	return ""
}

func (d Builder) execute(commandName string, commandArgs ...string) (builders.ArtifactPath, error) {
	args := append([]string{commandName}, commandArgs...)
	cmd := exec.Command("docker", args...)
	// TODO : Quand même passer par les pipes pour ne pas avoir l'effet "rien ne se passe" pendant les builds longs
	output, err := cmd.CombinedOutput()
	log.Printf("Docker command output: %v", string(output))
	if err != nil {
		log.Printf("Docker command error: %v", err)
		return "", err
	}
	return extractArtifact(string(output)), err
}

// Build Build the project artifact as a Docker image
func (d Builder) Build(contextPath string) (builders.ArtifactPath, error) {
	return d.execute("build", "-t", d.artifactName, contextPath)
}

// NewBuilder Create a new Docker builder
func NewBuilder(artifactName string) builders.Builder {
	return &Builder{
		artifactName,
	}
}
