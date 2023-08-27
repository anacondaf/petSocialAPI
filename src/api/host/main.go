package api

import (
	"context"
	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/mediatR"
	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/service"
	"net/http"

	"github.com/anacondaf/petSocialAPI/src/api/host/controller"
	"github.com/anacondaf/petSocialAPI/src/api/host/domain"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func apiVersioning(e *echo.Echo) *echo.Group {
	v1 := e.Group("/api/v1")

	return v1
}

func StartHttpServer(lc fx.Lifecycle, db *domain.Queries, logger *zap.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	v1 := apiVersioning(e)
	AddController(v1, db)

	mediatR.RegisterHandler(db, logger)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Start echo server at localhost:8000")

			// Start server
			go func() {
				if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
					e.Logger.Fatal("shutting down the server")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	return e
}

func AddController(g *echo.Group, db *domain.Queries) {
	controller.NewPetController(g, *service.NewPetService(db))
}
