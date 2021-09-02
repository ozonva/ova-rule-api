package configs

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

var ServerConfig *Server
var DatabaseConfig *Database
var KafkaConfig *Kafka

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	DBName       string `yaml:"dbname"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"user"`
	Password     string `yaml:"password"`
	PoolMaxConns int    `yaml:"pool_max_conns"`
}

type Kafka struct {
	Brokers []string `yaml:"brokers"`
}

func Load() {
	ServerConfig = &Server{}
	DatabaseConfig = &Database{}
	KafkaConfig = &Kafka{}

	mapper := map[string]interface{}{
		"database": DatabaseConfig,
		"server":   ServerConfig,
		"kafka": KafkaConfig,
	}

	configDir := getConfigDir()
	for name, config := range mapper {
		configFile := path.Join(configDir, name+".yml")
		loadConfigFromFile(configFile, config)
	}
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

func loadConfigFromFile(name string, config interface{}) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) GetAddress() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
