package metrics

import (
	"fmt"
	"net/http"

	"github.com/ozonva/ova-rule-api/configs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.Handle("/metrics/", promhttp.Handler())

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", configs.PrometheusConfig.Host, configs.PrometheusConfig.Port),
		Handler: mux,
	}

	log.Info().Msg("Запускаем сервер метрик...")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal().Err(err)
		}
	}()
}
