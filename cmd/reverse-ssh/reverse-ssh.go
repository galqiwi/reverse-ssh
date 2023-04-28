package main

import (
	"fmt"
	"github.com/galqiwi/reverse-ssh/internal/connector"
	"github.com/juju/fslock"
	"log"
)

func Main() error {
	args, err := getArgs()
	if err != nil {
		return err
	}

	lock := fslock.New(args.lockFile)

	lockErr := lock.TryLock()
	if lockErr != nil {
		return fmt.Errorf("falied to acquire lock > " + lockErr.Error())
	}
	defer func() {
		_ = lock.Unlock()
	}()

	return connector.Connect(args.connectionConfig)
}

func main() {
	err := Main()
	if err != nil {
		log.Fatal(err)
	}
}
