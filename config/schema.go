package config

import "time"

// ─── SCHEMA ────────────────────────────────────────────────────

// Config is the root of all configuration.
type Config struct {
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	SMTP     SMTP     `mapstructure:"smtp"`
	Redis    Redis    `mapstructure:"redis"`
	JWT      JWT      `mapstructure:"jwt"`
}

// App holds generic application settings.
type App struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
}

// Server holds HTTP server timeouts.
type Server struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
}

// Database holds all Postgres connection settings.
type Database struct {
	Driver          string        `mapstructure:"driver"`
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	Name            string        `mapstructure:"name"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	SSLMode         string        `mapstructure:"sslmode"`
	TimeZone        string        `mapstructure:"timezone"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// SMTP holds your mail server settings.
type SMTP struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// Logger holds your logger settings.
type Logger struct {
	Level            string    `mapstructure:"level"`
	encoding         string    `mapstructure:"format"`
	Development      bool      `mapstructure:"development"`
	LogDir           string    `mapstructure:"log_dir"`
	OutputPaths      []string  `mapstructure:"output_paths"`
	ErrorOutputPaths []string  `mapstructure:"error_output_paths"`
	Sampling         *Sampling `mapstructure:"sampling"`
}

// Sampling holds your logger sampling settings.
type Sampling struct {
	Initial    int `mapstructure:"initial"`
	Thereafter int `mapstructure:"thereafter"`
}

// Redis holds your cache server settings.
type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// JWT holds PASETO/JWT settings.
type JWT struct {
	PrivateKeyPath string        `mapstructure:"-"`
	PublicKeyPath  string        `mapstructure:"-"`
	Issuer         string        `mapstructure:"issuer"`
	AccessTTL      time.Duration `mapstructure:"access_ttl"`
	RefreshTTL     time.Duration `mapstructure:"refresh_ttl"`
}
