package wizard

import (
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/reverse_ssh"
)

func SetupSSHClient(cfg *reverse_ssh.ReverseSSHConfig) error {
	connectionConfig := cfg.ConnectionConfig
	var err error

	fmt.Print("SSH client setup\n\n")

	err = RequestConfigValue("path to ssh credentials", &connectionConfig.CredentialFile)
	if err != nil {
		return err
	}

	fmt.Println("\nDONE")

	return nil
}
