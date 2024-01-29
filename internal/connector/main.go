package connector

func Connect(config ConnectionConfig) error {
	for {
		err := RunSession(config)
		if err != nil {
			return err
		}
	}
}
