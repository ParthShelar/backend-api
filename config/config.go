package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

// SecretData holds the configuration data retrieved from Secret Manager.
type SecretData struct {
	ProjectID  string `json:"PROJECT_ID"`
	InstanceID string `json:"INSTANCE_ID"`
	DatabaseID string `json:"DATABASE_ID"`
	Port       string
}

// getSecret fetches the secret value from Google Cloud Secret Manager.
func getSecret(secretName string, ctx context.Context) (SecretData, error) {
	// Create a Secret Manager client.
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return SecretData{}, fmt.Errorf("failed to create secret manager client: %w", err)
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	// Access the secret version.
	resp, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return SecretData{}, fmt.Errorf("failed to access secret version: %w", err)
	}

	// Decode the secret payload.
	var secret SecretData
	err = json.Unmarshal(resp.Payload.Data, &secret)
	if err != nil {
		return SecretData{}, fmt.Errorf("failed to unmarshal secret payload: %w", err)
	}

	return secret, nil
}

// LoadConfig loads the configuration from Secret Manager.
func LoadConfig(ctx context.Context) *SecretData {
	secret, err := getSecret("projects/5987277861/secrets/db-config/versions/1", ctx)
	if err != nil {
		log.Fatalf("Failed to load secrets: %v", err)
	}

	// Return the secret data with the Port value set.
	return &SecretData{
		ProjectID:  secret.ProjectID,
		InstanceID: secret.InstanceID,
		DatabaseID: secret.DatabaseID,
		Port:       getEnv("PORT", "1000"),
	}
}

// GetDBPath constructs the database path for the SecretData.
func (s *SecretData) GetDBPath() string {
	return fmt.Sprintf("projects/%s/instances/%s/databases/%s", s.ProjectID, s.InstanceID, s.DatabaseID)
}

// getEnv retrieves the value of the environment variable or returns a fallback value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Using default value for %s: %s", key, fallback)
	return fallback
}
