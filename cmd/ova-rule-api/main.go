package main

import (
	"context"
	"github.com/ozonva/ova-rule-api/internal/flusher"
	"github.com/ozonva/ova-rule-api/internal/saver"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-rule-api/configs"
	api "github.com/ozonva/ova-rule-api/internal/app/ova_rule_api"
	"github.com/ozonva/ova-rule-api/internal/db"
	"github.com/ozonva/ova-rule-api/internal/repo"
)

func main() {
	ctx := context.Background()

	log.Info().Msg("Подгружаем конфиги...")
	configs.Load()

	log.Info().Msg("Подключаемся к базе...")
	pool, err := db.Connect(ctx, configs.DatabaseConfig)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer pool.Close()

	repo_ := repo.NewRepo(ctx, pool)
	flusher_ := flusher.NewFlusher(10, repo_)
	saver_ := saver.NewSaver(100, flusher_, time.Second)
	saver_.Init()

	log.Info().Msgf("Запускаем gRPC сервер: %s", configs.ServerConfig.GetAddress())
	apiServer := api.NewAPIServer(repo_, saver_)
	if err = api.Run(&apiServer); err != nil {
		log.Fatal().Err(err)
	}
}
