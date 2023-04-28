package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/connector"
	"os/user"
	"path/filepath"
	"time"
)

type Args struct {
	lockFile         string
	connectionConfig connector.ConnectionConfig
}

func getArgs() (*Args, error) {
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get user, %w", err)
	}

	lockFile := flag.String("lock-file", "/tmp/reverse-ssh-client-lock", "path of lock file")
	CredentialFile := flag.String("credential-file", filepath.Join(u.HomeDir, ".ssh/id_rsa"), "path to credentials")
	RemotePort := flag.Int("remote-port", 2222, "ssh tunnel port")
	LocalPort := flag.Int("local-port", 22, "port for localhost")
	HubPort := flag.Int("hub-port", 22, "hub ssh port")
	HubHostname := flag.String("hub-hostname", "", "hub ssh host")
	HubUsername := flag.String("hub-username", "", "hub")
	LocalUsername := flag.String("local-username", u.Username, "local username for healthchecks")
	HealthcheckCooldownSecs := flag.Int("healthcheck-cooldown-secs", 10, "halthcheck cooldown")
	HealthcheckTimeoutSecs := flag.Int("healthcheck-timeout-secs", 10, "halthcheck timeout")

	flag.Parse()

	if *HubHostname == "" {
		return nil, errors.New("hub-hostname is not set")
	}
	return &Args{
		*lockFile,
		connector.ConnectionConfig{
			CredentialFile:      *CredentialFile,
			RemotePort:          *RemotePort,
			LocalPort:           *LocalPort,
			HubPort:             *HubPort,
			HubHostname:         *HubHostname,
			HubUsername:         *HubUsername,
			LocalUsername:       *LocalUsername,
			HealthcheckCooldown: time.Second * time.Duration(*HealthcheckCooldownSecs),
			HealthcheckTimeout:  time.Second * time.Duration(*HealthcheckTimeoutSecs),
		},
	}, nil
}
