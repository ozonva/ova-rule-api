package main

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-rule-api/configs"
	api "github.com/ozonva/ova-rule-api/internal/app/ova_rule_api"
)

func main() {
	log.Info().Msg("Подгружаем конфиги...")
	configs.Load()

	log.Info().Msg(fmt.Sprintf("Запускаем gRPC сервер: %s%s", configs.ServerConfig.Host, configs.ServerConfig.Port))

	if err := api.Run(); err != nil {
		log.Fatal().Err(err)
	}
}
