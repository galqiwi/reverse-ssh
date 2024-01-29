package wizard

import (
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/reverse_ssh"
	"os"
	"strings"
)

func getCurrentExecutablePath() (string, error) {
	return os.Readlink("/proc/self/exe")
}

func getCommand(config *reverse_ssh.ReverseSSHConfig) (string, error) {
	connectionConfig := config.ConnectionConfig

	cmdPath, err := getCurrentExecutablePath()
	if err != nil {
		return "", err
	}

	cmd := []string{cmdPath, "connect"}
	cmd = append(cmd, "--credential-file", connectionConfig.CredentialFile)
	cmd = append(cmd, "--hub-hostname", connectionConfig.HubHostname)
	cmd = append(cmd, "--hub-port", fmt.Sprint(connectionConfig.HubPort))
	cmd = append(cmd, "--hub-username", connectionConfig.HubUsername)
	cmd = append(cmd, "--local-port", fmt.Sprint(connectionConfig.LocalPort))
	cmd = append(cmd, "--local-username", connectionConfig.LocalUsername)
	cmd = append(cmd, "--lock-file", config.LockFile)
	cmd = append(cmd, "--remote-port", fmt.Sprint(connectionConfig.RemotePort))

	output := strings.Builder{}
	for argIdx, arg := range cmd {
		output.WriteString(arg)
		if argIdx == len(cmd)-1 {
			continue
		}
		output.WriteString(" ")
	}
	return output.String(), nil
}
