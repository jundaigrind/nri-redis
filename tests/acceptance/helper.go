package acceptance

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func dockerComposeRunMode(vars []string, ports []string, container string, detached bool) (string, string, error) {
	cmdLine := []string{"run"}
	if detached {
		cmdLine = append(cmdLine, "-d")
	}
	cmdLine = append(cmdLine, "--name")
	cmdLine = append(cmdLine, container)
	for i := range vars {
		cmdLine = append(cmdLine, "-e")
		cmdLine = append(cmdLine, vars[i])
	}
	for p := range ports {
		cmdLine = append(cmdLine, fmt.Sprintf("-p%s", ports[p]))
	}
	cmdLine = append(cmdLine, container)
	fmt.Printf("executing: docker-compose %s\n", strings.Join(cmdLine, " "))
	cmd := exec.Command("docker-compose", cmdLine...)
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	stdout := outBuffer.String()
	stderr := errBuffer.String()
	return stdout, stderr, err
}

func dockerComposeRun(vars []string, container string) (string, string, error) {
	return dockerComposeRunMode(vars, []string{}, container, false)
}
