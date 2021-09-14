package configs

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var Config *CompositeConfig

type CompositeConfig struct {
	Server     *ServerConfig
	Database   *DatabaseConfig
	Kafka      *KafkaConfig
	Jaeger     *JaegerConfig
	Prometheus *PrometheusConfig
}

type ServerConfig struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	FlusherChunkSize int    `yaml:"flusher_chunk_size" mapstructure:"flusher_chunk_size"`
	SaverCapacity    uint   `yaml:"saver_capacity" mapstructure:"saver_capacity"`
}

type DatabaseConfig struct {
	DBName       string `yaml:"dbname"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"user" mapstructure:"user"`
	Password     string `yaml:"password"`
	PoolMaxConns int    `yaml:"pool_max_conns" mapstructure:"pool_max_conns"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
}

type JaegerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type PrometheusConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func Load() {
	setupViper()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	Config = &CompositeConfig{}
	viper.Unmarshal(Config)
}

func getConfigDir() string {
	configDir := os.Getenv("CONFIG_DIR")
	if configDir == "" {
		configDir = "development"
	}

	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	return path.Join(basePath, configDir)
}

func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func setupViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(getConfigDir())
}
