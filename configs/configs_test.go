package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigsLoad(t *testing.T) {
	serverCfg := Server{
		Host: "localhost",
		Port: "8000",
	}
	databaseCfg := Database{
		Host:     "localhost",
		Port:     "5432",
		Username: "ova",
		Password: "",
	}

	Load()

	assert.Equal(t, serverCfg, *ServerConfig)
	assert.Equal(t, databaseCfg, *DatabaseConfig)
}
