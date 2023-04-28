package main

import (
	"errors"
	"flag"
	"github.com/galqiwi/reverse-ssh/internal/connector"
	"os"
	"path/filepath"
	"time"
)

type Args struct {
	lockFile         string
	connectionConfig connector.ConnectionConfig
}

func getDefaultCredentialFile() (string, error) {
	user := os.Getenv("USER")
	if user == "" {
		return "", errors.New("failed to get $USER")
	}
	return filepath.Join("/home/", user, ".ssh/id_rsa"), nil
}

func getArgs() (*Args, error) {
	defaultCredentialFile, err := getDefaultCredentialFile()
	if err != nil {
		return nil, err
	}

	lockFile := flag.String("lock-file", "/tmp/reverse-ssh-client-lock", "path of lock file")
	CredentialFile := flag.String("credential-file", defaultCredentialFile, "path to credentials")
	RemotePort := flag.Int("remote-port", 2222, "ssh tunnel port")
	LocalPort := flag.Int("local-port", 22, "port for localhost")
	HubPort := flag.Int("hub-port", 22, "hub ssh port")
	HubHostname := flag.String("hub-hostname", "", "hub ssh host")
	HubUsername := flag.String("hub-username", "", "hub")
	LocalUsername := flag.String("local-username", os.Getenv("USER"), "local username for healthchecks")
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
