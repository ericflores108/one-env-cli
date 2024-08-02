package providermanager

type PostmanEnvironmentType string

type PostmanEnvironmentVariable struct {
	Key     string                 `json:"key"`
	Value   string                 `json:"value"`
	Enabled bool                   `json:"enabled"`
	Type    PostmanEnvironmentType `json:"type"`
}

type PostmanEnvironmentData struct {
	Name   string                       `json:"name"`
	Values []PostmanEnvironmentVariable `json:"values"`
}

type ProviderManager interface {
	// Provider - Get item to send to Plugin
	GetItem() error
	// Provider - If Plugin needs credentials, such as api-key, get from provider
	GetSecret() (string, error)
	// Provider - Translate Provider.Item to Plugin
	PostmanEnv() (*PostmanEnvironmentData, error)
}
