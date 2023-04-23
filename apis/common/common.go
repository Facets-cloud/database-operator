package common

import (
	"database/sql"
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

const (
	// ResourceCredentialsSecretEndpointKey is the key inside a connection secret for the connection endpoint
	ResourceCredentialsSecretEndpointKey = "endpoint"
	// ResourceCredentialsSecretPortKey is the key inside a connection secret for the connection port
	ResourceCredentialsSecretPortKey = "port"
	// ResourceCredentialsSecretUserKey is the key inside a connection secret for the connection user
	ResourceCredentialsSecretUserKey = "username"
	// ResourceCredentialsSecretPasswordKey is the key inside a connection secret for the connection password
	ResourceCredentialsSecretPasswordKey = "password"
	// ResourceCredentialsSecretDatabaseKey is the key inside a connection secret for the connection database
	ResourceCredentialsSecretDatabaseKey = "database"
	// ResourceCredentialsSecretDatabaseKey is the key inside a connection secret for the connection database
	ResourceCredentialsSecretSSLModeKey = "sslmode"

	// Status Type
	CREATE = "Create"
	SYNC   = "Sync"
	FAIL   = "Fail"
	DELETE = "Delete"
)

// A SecretReference is a reference to a secret in an arbitrary namespace.
type SecretReference struct {
	// Name of the secret.
	Name string `json:"name"`

	// Namespace of the secret.
	Namespace string `json:"namespace"`
}

// A SecretKeySelector is a reference to a secret key in an arbitrary namespace.
type SecretKeySelector struct {
	SecretReference `json:",inline"`

	// The key to select.
	Key string `json:"key"`
}

func ConnectToPostgres(connectionSecret *corev1.Secret) (*sql.DB, error) {
	endpoint := string(connectionSecret.Data[ResourceCredentialsSecretEndpointKey])
	port := string(connectionSecret.Data[ResourceCredentialsSecretPortKey])
	username := string(connectionSecret.Data[ResourceCredentialsSecretUserKey])
	password := string(connectionSecret.Data[ResourceCredentialsSecretPasswordKey])
	defaultDatabase := string(connectionSecret.Data[ResourceCredentialsSecretDatabaseKey])
	sslMode := string(connectionSecret.Data[ResourceCredentialsSecretSSLModeKey])

	if !(len(defaultDatabase) > 0) {
		defaultDatabase = "postgres"
	}

	if !(len(sslMode) > 0) {
		sslMode = "disable"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		endpoint,
		port,
		username,
		password,
		defaultDatabase,
		sslMode,
	)

	db, err := sql.Open("postgres", psqlInfo)

	return db, err
}