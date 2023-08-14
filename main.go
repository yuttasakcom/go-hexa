package main

import (
	"log"
	"os"

	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/yuttasakcom/go-hexa/src/core/common"
	"github.com/yuttasakcom/go-hexa/src/core/config"
	"github.com/yuttasakcom/go-hexa/src/core/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	// Liveness Probe
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	cfg := config.NewConfig(common.EnvFile())
	initLogger(cfg.App())
	initTracer(cfg)

	server.NewServer(cfg).Start()
	// @TODO: start worker
	// @TODO: start gRPC server
}

func initLogger(cfg config.App) {
	var level slog.LogLevel = slog.LevelInfo
	if cfg.DebugLog {
		level = slog.LevelDebug
	}

	config := slog.NewProductionConfig()
	config.LogLevel = level
	config.AppName = cfg.AppName
	config.Version = cfg.AppVersion

	slog.L().Configure(config)
}

func initTracer(cfg config.Configer) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		slog.L().Fatal("failed to create the Jaeger exporter: %v", slog.Error(err))
	}

	r, err := resource.Merge(resource.Default(), resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cfg.App().AppName),
		semconv.ServiceVersionKey.String(cfg.App().AppVersion),
		attribute.String("environment", cfg.App().AppEnv),
	))

	if err != nil {
		slog.L().Fatal("Error init Jaeger resource", slog.Error(err))
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(r),
	)

	otel.SetTracerProvider(tp)
}
