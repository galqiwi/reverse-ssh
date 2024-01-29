package sshutils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
)

type PortForwarding struct {
	LocalPort  int
	RemotePort int
}

type SSHArgs struct {
	CredentialsFile          string
	RemoteHost               string
	RemoteUsername           string
	RemotePort               int
	RemoteCommand            string
	HoldConnection           bool
	RemoteToLocalForwardings []PortForwarding
}

func RunSSH(ctx context.Context, args SSHArgs) error {
	cmdArgs := make([]string, 0)
	cmdArgs = append(cmdArgs, "-g")
	cmdArgs = append(cmdArgs, "-o", "StrictHostKeyChecking=no")
	cmdArgs = append(cmdArgs, "-o", "BatchMode=yes")
	cmdArgs = append(cmdArgs, "-i", args.CredentialsFile)
	cmdArgs = append(cmdArgs, "-p", fmt.Sprint(args.RemotePort))
	cmdArgs = append(cmdArgs, "-l", args.RemoteUsername)
	cmdArgs = append(cmdArgs, fmt.Sprint(args.RemoteHost))
	if args.RemoteCommand != "" {
		cmdArgs = append(cmdArgs, args.RemoteCommand)
	}
	for _, forwarding := range args.RemoteToLocalForwardings {
		cmdArgs = append(cmdArgs, "-R", fmt.Sprintf("%v:localhost:%v", forwarding.RemotePort, forwarding.LocalPort))
	}

	if args.HoldConnection && args.RemoteCommand != "" {
		return errors.New("HoldConnection and RemoteCommand cannot be set at the same time")
	}
	if args.HoldConnection {
		cmdArgs = append(cmdArgs, "-N")
	}

	cmd := exec.CommandContext(ctx, "ssh", cmdArgs...)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf(
			"%v%v%v",
			stdout.String(),
			stderr.String(),
			err.Error(),
		)
	}

	return nil
}
