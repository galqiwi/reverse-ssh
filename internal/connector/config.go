package connector

type ConnectionConfig struct {
	CredentialFile          string // done
	RemotePort              int    // done
	LocalPort               int    // done
	HubPort                 int    // done
	HubHostname             string // done
	HubUsername             string // done
	LocalUsername           string // done
	HealthcheckCooldownSecs int
	HealthcheckTimeoutSecs  int
}
