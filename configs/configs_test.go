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
		DBName:       "ova",
		Host:         "localhost",
		Port:         "5432",
		Username:     "ova",
		Password:     "iloveozon",
		PoolMaxConns: 10,
	}

	Load()

	assert.Equal(t, serverCfg, *ServerConfig)
	assert.Equal(t, databaseCfg, *DatabaseConfig)
}
