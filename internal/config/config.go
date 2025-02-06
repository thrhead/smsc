package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Database   DatabaseConfig   `mapstructure:"database"`
	SMPP       SMPPConfig       `mapstructure:"smpp"`
	Sigtran    SigtranConfig    `mapstructure:"sigtran"`
	Security   SecurityConfig   `mapstructure:"security"`
	Routing    RoutingConfig    `mapstructure:"routing"`
	Monitoring MonitoringConfig `mapstructure:"monitoring"`
	Logging    LoggingConfig    `mapstructure:"logging"`
	Queue      QueueConfig      `mapstructure:"queue"`
	RateLimit  RateLimitConfig  `mapstructure:"rate_limiting"`
}

type ServerConfig struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type SMPPConfig struct {
	Host      string        `mapstructure:"host"`
	Port      int          `mapstructure:"port"`
	TLSPort   int          `mapstructure:"tls_port"`
	SystemID  string        `mapstructure:"system_id"`
	Password  string        `mapstructure:"password"`
	Timeout   time.Duration `mapstructure:"timeout"`
}

type SigtranConfig struct {
	SCTP  SCTPConfig  `mapstructure:"sctp"`
	M3UA  M3UAConfig  `mapstructure:"m3ua"`
	SCCP  SCCPConfig  `mapstructure:"sccp"`
}

type SCTPConfig struct {
	LocalPort  int `mapstructure:"local_port"`
	RemotePort int `mapstructure:"remote_port"`
	MaxStreams int `mapstructure:"max_streams"`
}

type M3UAConfig struct {
	LocalPointCode    int `mapstructure:"local_point_code"`
	RemotePointCode   int `mapstructure:"remote_point_code"`
	NetworkIndicator  int `mapstructure:"network_indicator"`
	RoutingContext   int `mapstructure:"routing_context"`
}

type SCCPConfig struct {
	LocalGT          string `mapstructure:"local_gt"`
	TranslationType  int    `mapstructure:"translation_type"`
}

type SecurityConfig struct {
	JWTSecret   string        `mapstructure:"jwt_secret"`
	TokenExpiry time.Duration `mapstructure:"token_expiry"`
	TLSCert     string        `mapstructure:"tls_cert"`
	TLSKey      string        `mapstructure:"tls_key"`
}

type RoutingConfig struct {
	DefaultRoute   string           `mapstructure:"default_route"`
	MaxRetries    int              `mapstructure:"max_retries"`
	RetryInterval time.Duration    `mapstructure:"retry_interval"`
	Operators     []OperatorConfig `mapstructure:"operators"`
}

type OperatorConfig struct {
	Name     string `mapstructure:"name"`
	Priority int    `mapstructure:"priority"`
	Weight   int    `mapstructure:"weight"`
	MaxTPS   int    `mapstructure:"max_tps"`
}

type MonitoringConfig struct {
	PrometheusEnabled  bool          `mapstructure:"prometheus_enabled"`
	MetricsPath       string        `mapstructure:"metrics_path"`
	CollectionInterval time.Duration `mapstructure:"collection_interval"`
}

type LoggingConfig struct {
	Level    string `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Output   string `mapstructure:"output"`
	FilePath string `mapstructure:"file_path"`
}

type QueueConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type RateLimitConfig struct {
	Enabled           bool `mapstructure:"enabled"`
	RequestsPerSecond int  `mapstructure:"requests_per_second"`
	Burst            int  `mapstructure:"burst"`
}

// Load loads the configuration from a file
func Load(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Server.Port <= 0 {
		return fmt.Errorf("invalid server port")
	}

	if c.Database.Port <= 0 {
		return fmt.Errorf("invalid database port")
	}

	if c.SMPP.Port <= 0 {
		return fmt.Errorf("invalid SMPP port")
	}

	if c.Security.JWTSecret == "" {
		return fmt.Errorf("JWT secret is required")
	}

	return nil
} 