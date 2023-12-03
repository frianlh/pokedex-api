package migrations

import (
	"errors"
	"github.com/frianlh/pokedex-api/libs/constants"
	"os"
	"strconv"
)

// MigrationConfigInterface is
type MigrationConfigInterface interface {
	Read() (migrationConfig *MigrationConfig, err error)
}

// MigrationConfig is
type MigrationConfig struct {
	MigrationPath  string
	PostgresConfig postgresConfig
}

type postgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// NewMigrationConfig is
func NewMigrationConfig() MigrationConfig {
	return MigrationConfig{}
}

// Read is
func (mc MigrationConfig) Read() (migrationConfig *MigrationConfig, err error) {
	// postgres config
	postgresHostStr := os.Getenv("POSTGRES_DB_HOST")
	if postgresHostStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	mc.PostgresConfig.Host = postgresHostStr
	postgresPortStr := os.Getenv("POSTGRES_DB_PORT")
	if postgresPortStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	postgresPort, err := strconv.Atoi(postgresPortStr)
	if err == nil {
		mc.PostgresConfig.Port = postgresPort
	}
	postgresUserStr := os.Getenv("POSTGRES_DB_USER")
	if postgresUserStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	mc.PostgresConfig.User = postgresUserStr
	postgresPasswordStr := os.Getenv("POSTGRES_DB_PASSWORD")
	if postgresPasswordStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	mc.PostgresConfig.Password = postgresPasswordStr
	postgresDbNameStr := os.Getenv("POSTGRES_DB_NAME")
	if postgresDbNameStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	mc.PostgresConfig.DbName = postgresDbNameStr

	// migration path config
	migrationPathStr := os.Getenv("MIGRATION_PATH")
	if migrationPathStr == "" {
		return nil, errors.New(constants.MigrationInvalidEnv)
	}
	mc.MigrationPath = migrationPathStr

	return &mc, nil
}
