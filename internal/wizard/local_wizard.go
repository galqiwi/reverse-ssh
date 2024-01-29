package wizard

import (
	"context"
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/reverse_ssh"
	"github.com/galqiwi/reverse-ssh/internal/sshutils"
	"time"
)

func SetupLocalConnectionParams(cfg *reverse_ssh.ReverseSSHConfig) error {
	connectionConfig := cfg.ConnectionConfig
	var err error

	fmt.Print("Local connection setup\n\n")

	err = RequestConfigValue("username for health checks", &connectionConfig.LocalUsername)
	if err != nil {
		return err
	}

	err = RequestConfigIntValue("local ssh server port", &connectionConfig.LocalPort)
	if err != nil {
		return err
	}

	fmt.Print("\nChecking if health check is possible...  ")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = sshutils.RunSSH(ctx, sshutils.SSHArgs{
		CredentialsFile:          connectionConfig.CredentialFile,
		RemoteHost:               "localhost",
		RemoteUsername:           connectionConfig.LocalUsername,
		RemotePort:               connectionConfig.LocalPort,
		RemoteCommand:            "echo ok",
		HoldConnection:           false,
		RemoteToLocalForwardings: nil,
	})

	if err != nil {
		fmt.Println("\033[91mFAIL\033[0m")
	} else {
		fmt.Println("\033[92mOK\033[0m")
	}

	if err != nil {
		return err
	}

	fmt.Println("DONE")

	return nil
}
