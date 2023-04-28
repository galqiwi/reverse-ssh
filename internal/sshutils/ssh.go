package sshutils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

func beautifyCmdArgs(args []string) string {
	output := make([]string, 0, len(args))

	for _, argument := range args {
		if strings.Contains(argument, " ") {
			argument = fmt.Sprintf("'%s'", argument)
		}
		output = append(output, argument)
	}

	return strings.Join(output, " ")
}

func RunSSH(ctx context.Context, args SSHArgs) error {
	cmdArgs := make([]string, 0)
	cmdArgs = append(cmdArgs, "-g")
	cmdArgs = append(cmdArgs, "-o", "StrictHostKeyChecking=no")
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

	log.Println("running ssh", beautifyCmdArgs(cmdArgs))
	cmd := exec.CommandContext(ctx, "ssh", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
