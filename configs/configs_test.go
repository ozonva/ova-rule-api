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
	kafkaCfg := Kafka{
		Brokers: []string{"127.0.0.1:9092"},
	}
	jaegerCfg := Jaeger{
		Host: "localhost",
		Port: "6831",
	}


	Load()

	assert.Equal(t, serverCfg, *ServerConfig)
	assert.Equal(t, databaseCfg, *DatabaseConfig)
	assert.Equal(t, kafkaCfg, *KafkaConfig)
	assert.Equal(t, jaegerCfg, *JaegerConfig)
}
