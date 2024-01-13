package env

import (
	"github.com/romanfomindev/in-memory-db/internal/config"
	"os"
)

const AppEnvDev = "dev"

type AppConfig struct {
	appEnv string
}

func NewAppConfig() config.AppConfig {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = AppEnvDev
	}

	return &AppConfig{
		appEnv: env,
	}
}

func (cfg *AppConfig) Env() string {
	return cfg.appEnv
}
