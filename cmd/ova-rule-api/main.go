package main

import (
	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-rule-api/configs"
	api "github.com/ozonva/ova-rule-api/internal/app/ova_rule_api"
)

func main() {
	log.Info().Msg("Подгружаем конфиги...")
	configs.Load()

	log.Info().Msgf("Запускаем gRPC сервер: %s", configs.ServerConfig.GetAddress())

	if err := api.Run(); err != nil {
		log.Fatal().Err(err)
	}
}
