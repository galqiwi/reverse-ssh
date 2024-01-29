package connector

type ConnectionConfig struct {
	CredentialFile          string
	RemotePort              int
	LocalPort               int
	HubPort                 int
	HubHostname             string
	HubUsername             string
	LocalUsername           string
	HealthcheckCooldownSecs int
	HealthcheckTimeoutSecs  int
}
