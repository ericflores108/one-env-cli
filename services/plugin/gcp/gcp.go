package gcp

import (
	"github.com/ericflores108/one-env-cli/providermanager"
)

type PluginJob struct {
	Provider providermanager.ProviderManager
}

// createSecret creates a new secret with the given name. A secret is a logical
// wrapper around a collection of secret versions. Secret versions hold the
// actual secret material.
// func CreateSecret(w io.Writer, parent, id string) error {
// 	// parent := "projects/my-project"
// 	// id := "my-secret"

// 	// Create the client.
// 	ctx := context.Background()
// 	client, err := secretmanager.NewClient(ctx)
// 	if err != nil {
// 		return fmt.Errorf("failed to create secretmanager client: %w", err)
// 	}
// 	defer client.Close()

// 	// Build the request.
// 	req := &secretmanagerpb.CreateSecretRequest{
// 		Parent:   parent,
// 		SecretId: id,
// 		Secret: &secretmanagerpb.Secret{
// 			Replication: &secretmanagerpb.Replication{
// 				Replication: &secretmanagerpb.Replication_Automatic_{
// 					Automatic: &secretmanagerpb.Replication_Automatic{},
// 				},
// 			},
// 		},
// 	}

// 	// Call the API.
// 	result, err := client.CreateSecret(ctx, req)
// 	if err != nil {
// 		return fmt.Errorf("failed to create secret: %w", err)
// 	}
// 	fmt.Fprintf(w, "Created secret: %s\n", result.Name)
// 	return nil
// }

func NewPluginJob(providerManager providermanager.ProviderManager) *PluginJob {
	return &PluginJob{
		Provider: providerManager,
	}
}
