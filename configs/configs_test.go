package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigsLoad(t *testing.T) {
	t.Parallel()

	serverCfg := Server{
		Host:             "localhost",
		Port:             "8000",
		FlusherChunkSize: 10,
		SaverCapacity:    100,
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
	prometheusCfg := Prometheus{
		Host: "localhost",
		Port: "9102",
	}

	Load()

	assert.Equal(t, serverCfg, *ServerConfig)
	assert.Equal(t, databaseCfg, *DatabaseConfig)
	assert.Equal(t, kafkaCfg, *KafkaConfig)
	assert.Equal(t, jaegerCfg, *JaegerConfig)
	assert.Equal(t, prometheusCfg, *PrometheusConfig)
}
