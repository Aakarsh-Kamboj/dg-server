package config

import (
	"fmt"
	"github.com/joho/godotenv"          // optional: for local dev
	"github.com/mitchellh/mapstructure" // for custom decode hooks
	"github.com/spf13/viper"
	"log"
	"strings"
)

// ─── LOADER ────────────────────────────────────────────────────

func Load() (*Config, error) {
	// 1. Load .env into process ENV (dev only)
	_ = godotenv.Load()

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 2. Bind every credential ENV to its key
	bindings := map[string]string{
		// Postgres
		"database.host":     "DB_HOST",
		"database.port":     "DB_PORT",
		"database.name":     "DB_NAME",
		"database.user":     "DB_USER",
		"database.password": "DB_PASSWORD",
		"database.sslmode":  "DB_SSLMODE",
		// SMTP
		"smtp.host":     "SMTP_HOST",
		"smtp.port":     "SMTP_PORT",
		"smtp.user":     "SMTP_USER",
		"smtp.password": "SMTP_PASSWORD",
		// Redis
		"redis.host":     "REDIS_HOST",
		"redis.port":     "REDIS_PORT",
		"redis.password": "REDIS_PASSWORD",
		// JWT
		"jwt.private_key": "JWT_PRIVATE_KEY_PATH",
		"jwt.public_key":  "JWT_PUBLIC_KEY_PATH",
	}

	for key, env := range bindings {
		if err := v.BindEnv(key, env); err != nil {
			return nil, fmt.Errorf("bind env %s: %w", env, err)
		}
	}

	// 3. Read the YAML defaults
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	// 4. Unmarshal into our schema, with a *composed* hook:
	var cfg Config
	decodeHook := mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(), // "15s" → time.Duration
		mapstructure.StringToSliceHookFunc(","),     // "a,b,c" → []string
	)
	if err := v.Unmarshal(
		&cfg,
		viper.DecodeHook(decodeHook),
	); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

// ProvideConfig for dependency injection
func ProvideConfig() *Config {
	cfg, err := Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	return cfg
}
