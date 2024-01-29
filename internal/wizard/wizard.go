package wizard

import (
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/reverse_ssh"
	"github.com/spf13/cobra"
)

var WizardCmd = &cobra.Command{
	Use: "wizard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Main()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Main() error {
	config, err := reverse_ssh.NewDefaultConfig()
	if err != nil {
		return err
	}

	err = SetupSSHClient(config)
	if err != nil {
		return err
	}

	PrintSplitter()

	err = SetupHubParams(config)
	if err != nil {
		return err
	}

	PrintSplitter()

	err = SetupLocalConnectionParams(config)
	if err != nil {
		return err
	}

	PrintSplitter()

	err = SetupReverseSSHParams(config)
	if err != nil {
		return err
	}

	PrintSplitter()

	err = SetupLockParams(config)
	if err != nil {
		return err
	}

	PrintSplitter()

	cmd, err := getCommand(config)
	if err != nil {
		return err
	}

	fmt.Printf("Command for reverse ssh connection:\n\n%v\n", cmd)

	return nil
}
