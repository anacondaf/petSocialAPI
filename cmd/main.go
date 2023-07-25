package main

import (
	"github.com/anacondaf/petSocialAPI/src/api/host/domain"
	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/logger"

	api "github.com/anacondaf/petSocialAPI/src/api/host"
	_ "github.com/go-sql-driver/mysql"
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
