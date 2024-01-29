package reverse_ssh

import (
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/connector"
	"os/user"
	"path/filepath"
)

type ReverseSSHConfig struct {
	ConnectionConfig connector.ConnectionConfig
	LockFile         string
}

func NewDefaultConfig() (*ReverseSSHConfig, error) {
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get user, %w", err)
	}

	connectionConfig := connector.ConnectionConfig{
		CredentialFile:          filepath.Join(u.HomeDir, ".ssh/id_rsa"),
		RemotePort:              2222,
		LocalPort:               22,
		HubPort:                 22,
		HubHostname:             "",
		HubUsername:             "hub",
		LocalUsername:           u.Username,
		HealthcheckCooldownSecs: 10,
		HealthcheckTimeoutSecs:  10,
	}

	return &ReverseSSHConfig{
		ConnectionConfig: connectionConfig,
		LockFile:         "/tmp/reverse-ssh-client-lock",
	}, nil
}
