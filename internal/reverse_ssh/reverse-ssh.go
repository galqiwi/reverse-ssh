package reverse_ssh

import (
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/connector"
	"github.com/juju/fslock"
	"github.com/spf13/cobra"
	"log"
)

var config *ReverseSSHConfig

var ConnectCmd = &cobra.Command{
	Use: "connect",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Main()
	},
}

func init() {
	var err error
	config, err = NewDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	connectionConfig := config.ConnectionConfig

	ConnectCmd.PersistentFlags().StringVar(&config.LockFile, "lock-file", config.LockFile, "path of lock file")

	ConnectCmd.PersistentFlags().StringVar(
		&connectionConfig.CredentialFile, "credential-file", connectionConfig.CredentialFile, "path to credentials",
	)
	ConnectCmd.PersistentFlags().IntVar(
		&connectionConfig.RemotePort, "remote-port", connectionConfig.RemotePort, "ssh tunnel port",
	)
	ConnectCmd.PersistentFlags().IntVar(
		&connectionConfig.LocalPort, "local-port", connectionConfig.LocalPort, "port for localhost",
	)
	ConnectCmd.PersistentFlags().IntVar(
		&connectionConfig.HubPort, "hub-port", connectionConfig.HubPort, "hub ssh port",
	)
	ConnectCmd.PersistentFlags().StringVar(
		&connectionConfig.HubHostname, "hub-hostname", connectionConfig.HubHostname, "hub ssh host",
	)
	ConnectCmd.PersistentFlags().StringVar(
		&connectionConfig.HubUsername, "hub-username", connectionConfig.HubUsername, "hub username",
	)
	ConnectCmd.PersistentFlags().StringVar(
		&connectionConfig.LocalUsername, "local-username", connectionConfig.LocalUsername, "local username for healthchecks",
	)
	ConnectCmd.PersistentFlags().IntVar(
		&connectionConfig.HealthcheckCooldownSecs, "healthcheck-cooldown-secs", connectionConfig.HealthcheckCooldownSecs, "healthcheck cooldown",
	)
	ConnectCmd.PersistentFlags().IntVar(
		&connectionConfig.HealthcheckTimeoutSecs, "healthcheck-timeout-secs", connectionConfig.HealthcheckTimeoutSecs, "healthcheck timeout",
	)
}

func Main() error {
	lock := fslock.New(config.LockFile)

	lockErr := lock.TryLock()
	if lockErr != nil {
		return fmt.Errorf("falied to acquire lock > " + lockErr.Error())
	}
	defer func() {
		_ = lock.Unlock()
	}()

	return connector.Connect(config.ConnectionConfig)
}
