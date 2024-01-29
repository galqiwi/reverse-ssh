package connector

import (
	"context"
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/sshutils"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"time"
)

func CheckConnection(config ConnectionConfig) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.HealthcheckTimeoutSecs)*time.Second)
	defer cancel()

	challenge := filepath.Join(os.TempDir(), uuid.New().String())
	defer func() {
		_ = os.Remove(challenge)
	}()

	err := sshutils.RunSSH(ctx, sshutils.SSHArgs{
		RemoteHost:      config.HubHostname,
		RemotePort:      config.RemotePort,
		RemoteUsername:  config.LocalUsername,
		CredentialsFile: config.CredentialFile,
		RemoteCommand:   fmt.Sprintf("touch %s", challenge),
	})
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(challenge); err == nil {
		return true, nil
	}
	return false, nil
}
