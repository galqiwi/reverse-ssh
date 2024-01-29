package wizard

import (
	"context"
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/connector"
	"github.com/galqiwi/reverse-ssh/internal/reverse_ssh"
	"time"
)

func SetupReverseSSHParams(cfg *reverse_ssh.ReverseSSHConfig) error {
	connectionConfig := cfg.ConnectionConfig
	var err error

	fmt.Print("Reverse ssh tunneling setup\n\n")

	err = RequestConfigIntValue("remote port", &connectionConfig.RemotePort)
	if err != nil {
		return err
	}

	fmt.Print("\nChecking reverse ssh tunneling...  ")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	go func() {
		connector.RunSession(ctx, *connectionConfig)
	}()

	time.Sleep(time.Second)

	ok, err := connector.CheckConnection(*connectionConfig)

	if err != nil || !ok {
		fmt.Println("\033[91mFAIL\033[0m")
	} else {
		fmt.Println("\033[92mOK\033[0m")
	}

	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("healthcheck failed")
	}

	fmt.Println("DONE")

	return nil
}

func SetupLockParams(cfg *reverse_ssh.ReverseSSHConfig) error {
	var err error

	fmt.Print("Lock setup\n\n")

	err = RequestConfigValue("lock path", &cfg.LockFile)
	if err != nil {
		return err
	}

	fmt.Println("DONE")

	return nil
}
