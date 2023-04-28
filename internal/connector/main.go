package connector

import "time"

type ConnectionConfig struct {
	CredentialFile      string
	RemotePort          int
	LocalPort           int
	HubPort             int
	HubHostname         string
	HubUsername         string
	LocalUsername       string
	HealthcheckCooldown time.Duration
	HealthcheckTimeout  time.Duration
}

func Connect(config ConnectionConfig) error {
	for {
		err := RunSession(config)
		if err != nil {
			return err
		}
	}
}
