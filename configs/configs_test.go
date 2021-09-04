package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigsLoad(t *testing.T) {
	t.Parallel()

	serverCfg := ServerConfig{
		Host:             "localhost",
		Port:             "8000",
		FlusherChunkSize: 10,
		SaverCapacity:    100,
	}
	databaseCfg := DatabaseConfig{
		DBName:       "ova",
		Host:         "localhost",
		Port:         "5432",
		Username:     "ova",
		Password:     "iloveozon",
		PoolMaxConns: 10,
	}
	kafkaCfg := KafkaConfig{
		Brokers: []string{"127.0.0.1:9092"},
	}
	jaegerCfg := JaegerConfig{
		Host: "localhost",
		Port: "6831",
	}
	prometheusCfg := PrometheusConfig{
		Host: "localhost",
		Port: "9102",
	}

	cfg := CompositeConfig{
		Server:     &serverCfg,
		Database:   &databaseCfg,
		Kafka:      &kafkaCfg,
		Jaeger:     &jaegerCfg,
		Prometheus: &prometheusCfg,
	}

	Load()

	assert.Equal(t, cfg, *Config)
}
