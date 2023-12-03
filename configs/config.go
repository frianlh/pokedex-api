package configs

import (
	"errors"
	"github.com/frianlh/pokedex-api/connections"
	"github.com/frianlh/pokedex-api/libs/constants"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

// ConfigInterface is
type ConfigInterface interface {
	Read() (*Config, error)
}

// Config is
type Config struct {
	PortApi        int
	BaseURL        string
	AppName        string
	AppEnv         string
	AppMode        string
	PostgresConfig postgresConfig
	JWTKey         string
	TimeoutCtx     time.Duration
}

type postgresConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	DbName      string
	DbConn      *gorm.DB
	MaxOpenConn int
	MaxIdleConn int
}

// NewConfig is
func NewConfig() ConfigInterface {
	timeoutCtx := time.Duration(30) * time.Second
	return &Config{
		TimeoutCtx: timeoutCtx,
	}
}

// Read is
func (c Config) Read() (config *Config, err error) {
	// port config
	c.PortApi = 8080
	portApiStr := os.Getenv("PORT_API")
	portApi, err := strconv.Atoi(portApiStr)
	if err == nil {
		c.PortApi = portApi
	}

	// base url config
	baseURLStr := os.Getenv("BASE_URL")
	if baseURLStr == "" {
		return nil, errors.New(constants.BaseURLInvalidEnv)
	}
	c.BaseURL = baseURLStr

	// app config
	appNameStr := os.Getenv("APP_NAME")
	if appNameStr == "" {
		return nil, errors.New(constants.AppInvalidEnv)
	}
	c.AppName = appNameStr
	appEnvStr := os.Getenv("APP_ENV")
	if appEnvStr == "" {
		return nil, errors.New(constants.AppInvalidEnv)
	}
	c.AppEnv = appEnvStr
	appModeStr := os.Getenv("APP_MODE")
	if appModeStr == "" {
		return nil, errors.New(constants.AppInvalidEnv)
	}
	c.AppMode = appModeStr

	// postgres config
	postgresHostStr := os.Getenv("POSTGRES_DB_HOST")
	if postgresHostStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	c.PostgresConfig.Host = postgresHostStr
	postgresPortStr := os.Getenv("POSTGRES_DB_PORT")
	if postgresPortStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	postgresPort, err := strconv.Atoi(postgresPortStr)
	if err == nil {
		c.PostgresConfig.Port = postgresPort
	}
	postgresUserStr := os.Getenv("POSTGRES_DB_USER")
	if postgresUserStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	c.PostgresConfig.User = postgresUserStr
	postgresPasswordStr := os.Getenv("POSTGRES_DB_PASSWORD")
	if postgresPasswordStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	c.PostgresConfig.Password = postgresPasswordStr
	postgresDbNameStr := os.Getenv("POSTGRES_DB_NAME")
	if postgresDbNameStr == "" {
		return nil, errors.New(constants.PostgresInvalidEnv)
	}
	c.PostgresConfig.DbName = postgresDbNameStr
	c.PostgresConfig.MaxOpenConn = 20
	postgresMaxOpenConnStr := os.Getenv("POSTGRES_DB_MAX_OPEN_CONN")
	if postgresMaxOpenConnStr != "" {
		postgresMaxOpenConn, err := strconv.Atoi(postgresMaxOpenConnStr)
		if err == nil {
			c.PostgresConfig.MaxOpenConn = postgresMaxOpenConn
		}
	}
	c.PostgresConfig.MaxIdleConn = 20
	postgresMaxIdleConnStr := os.Getenv("POSTGRES_DB_MAX_IDLE_CONN")
	if postgresMaxIdleConnStr != "" {
		postgresMaxIdleConn, err := strconv.Atoi(postgresMaxIdleConnStr)
		if err == nil {
			c.PostgresConfig.MaxIdleConn = postgresMaxIdleConn
		}
	}
	postgresConn, err := connections.PostgresConn(c.PostgresConfig.Host, c.PostgresConfig.User, c.PostgresConfig.Password, c.PostgresConfig.DbName, c.AppMode, c.PostgresConfig.Port, c.PostgresConfig.MaxOpenConn, c.PostgresConfig.MaxIdleConn)
	if err != nil {
		return nil, err
	}
	c.PostgresConfig.DbConn = postgresConn

	// JWT key config
	jwtKeyStr := os.Getenv("JWT_KEY")
	if jwtKeyStr == "" {
		return nil, errors.New(constants.JWTKeyEnv)
	}
	c.JWTKey = jwtKeyStr

	return &c, nil
}
