package providermanager

type ProviderManager interface {
	GetItem(name string) error
}
