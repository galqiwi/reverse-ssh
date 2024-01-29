package connector

func Connect(config ConnectionConfig) error {
	for {
		err := RunHealthCheckedSession(config)
		if err != nil {
			return err
		}
	}
}
