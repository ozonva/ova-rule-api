package main

import (
	"context"
	"time"

	"github.com/ozonva/ova-rule-api/internal/flusher"
	"github.com/ozonva/ova-rule-api/internal/kafka"
	"github.com/ozonva/ova-rule-api/internal/metrics"
	"github.com/ozonva/ova-rule-api/internal/saver"
	"github.com/ozonva/ova-rule-api/internal/tracer"

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

	producer, err := kafka.NewAsyncProducer(configs.KafkaConfig.Brokers)
	if err != nil {
		log.Fatal().Err(err)
	}

	closer, err := tracer.InitTracer()
	if err != nil {
		log.Fatal().Err(err)
	}
	defer closer.Close()

	ruleMetrics := metrics.NewMetrics()
	ruleRepo := repo.NewRepo(ctx, pool, producer)
	ruleFlusher := flusher.NewFlusher(configs.ServerConfig.FlusherChunkSize, ruleRepo)
	ruleSaver := saver.NewSaver(configs.ServerConfig.SaverCapacity, ruleFlusher, time.Second)
	ruleSaver.Init()

	metrics.RunServer()

	log.Info().Msgf("Запускаем gRPC сервер: %s", configs.ServerConfig.GetAddress())

	apiServer := api.NewAPIServer(ruleRepo, ruleSaver, ruleMetrics)
	if err = api.Run(&apiServer); err != nil {
		log.Fatal().Err(err)
	}
}
