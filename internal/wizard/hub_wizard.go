package wizard

import (
	"context"
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/reverse_ssh"
	"github.com/galqiwi/reverse-ssh/internal/sshutils"
	"time"
)

func SetupHubParams(cfg *reverse_ssh.ReverseSSHConfig) error {
	connectionConfig := cfg.ConnectionConfig
	var err error

	fmt.Print("Hub connection setup\n\n")

	err = RequestConfigValue("hub address", &connectionConfig.HubHostname)
	if err != nil {
		return err
	}

	err = RequestConfigValue("hub username", &connectionConfig.HubUsername)
	if err != nil {
		return err
	}

	err = RequestConfigIntValue("hub port", &connectionConfig.HubPort)
	if err != nil {
		return err
	}

	fmt.Print("\nChecking connection to the hub...  ")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = sshutils.RunSSH(ctx, sshutils.SSHArgs{
		CredentialsFile:          connectionConfig.CredentialFile,
		RemoteHost:               connectionConfig.HubHostname,
		RemoteUsername:           connectionConfig.HubUsername,
		RemotePort:               connectionConfig.HubPort,
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
