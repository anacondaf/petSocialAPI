package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go-echo/api"
	"go-echo/api/domain"
	"go-echo/api/infrastructure/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			domain.UseSqlc,
			logger.UseZap,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(api.StartHttpServer),
	).Run()
}
