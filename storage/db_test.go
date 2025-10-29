package storage

import (
	"context"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Helper to reset env between tests
func resetEnv() {
	os.Unsetenv("DB_DRIVER")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("SSL_MODE")
}

func TestInitDatabasePostgres(t *testing.T) {
	resetEnv()

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "messages",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %v", err)
	}
	defer container.Terminate(ctx)

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Failed to get mapped port: %v", err)
	}

	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", port.Port())
	os.Setenv("DB_USER", "test")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_NAME", "messages")
	os.Setenv("SSL_MODE", "disable")

	store, _ := InitDatabase()
	if store == nil {
		t.Fatal("Expected Postgres store, got nil")
	}
}

func TestInitDatabaseSQLite(t *testing.T) {
	resetEnv()
	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("DB_PATH", ":memory:")

	store, _ := InitDatabase()
	if store == nil {
		t.Fatal("Expected SQLite store, got nil")
	}
}

func TestInitDatabaseUnsupportedDriver(t *testing.T) {
	resetEnv()
	os.Setenv("DB_DRIVER", "oracle") // unsupported

	store, err := InitDatabase()
	if err == nil {
		t.Error("Expected error for unsupported driver, got nil")
	}
	if store != nil {
		t.Error("Expected nil store for unsupported driver")
	}
}

func TestInitDatabaseConnectionFailure(t *testing.T) {
	resetEnv()
	os.Setenv("DB_DRIVER", "postgres") // unsupported

	store, err := InitDatabase()
	if err == nil {
		t.Error("Expected error for unsupported driver, got nil")
	}
	if store != nil {
		t.Error("Expected nil store for unsupported driver")
	}
}
