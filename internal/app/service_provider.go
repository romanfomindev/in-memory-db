package app

import (
	"github.com/romanfomindev/in-memory-db/internal/config"
	"github.com/romanfomindev/in-memory-db/internal/config/env"
)

type ServiceProvider struct {
	appConfig config.AppConfig
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) AppConfig() config.AppConfig {
	if s.appConfig == nil {
		s.appConfig = env.NewAppConfig()
	}

	return s.appConfig
}
