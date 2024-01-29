package main

import (
	"github.com/galqiwi/reverse-ssh/internal/reverse_ssh"
	"github.com/spf13/cobra"
	"log"
)

var ReverseSSHCmd = &cobra.Command{
	Use:   "reverse-ssh",
	Short: "Reverse ssh tunneling script",
}

func init() {
	ReverseSSHCmd.AddCommand(reverse_ssh.ConnectCmd)
}

func main() {
	err := Main()
	if err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	return ReverseSSHCmd.Execute()
}
