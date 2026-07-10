package config_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_llm_pii/app/internal/config"
)

func TestGetEnvironmentConfig(t *testing.T) {
	envConfig := config.NewEnvironmentConfig()

	if envConfig == nil {
		t.Fatal("expected config, got nil")
	}

}
