package connector

import (
	"context"
	"github.com/galqiwi/reverse-ssh/internal/sshutils"
	"log"
	"sync"
	"time"
)

func RunSession(config ConnectionConfig) error {
	sessionContext, killSession := context.WithCancel(context.Background())
	defer killSession()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer killSession()
		_ = sshutils.RunSSH(sessionContext, sshutils.SSHArgs{
			RemoteHost:      config.HubHostname,
			RemotePort:      config.HubPort,
			RemoteUsername:  config.HubUsername,
			CredentialsFile: config.CredentialFile,
			HoldConnection:  true,
			RemoteToLocalForwardings: []sshutils.PortForwarding{{
				LocalPort:  config.LocalPort,
				RemotePort: config.RemotePort,
			}},
		})
	}()

	for {
		time.Sleep(config.HealthcheckCooldown)
		ok, err := CheckConnection(config)
		if err != nil || !ok {
			break
		}
	}

	log.Println("dropping connection")
	killSession()

	wg.Wait()
	return nil
}
