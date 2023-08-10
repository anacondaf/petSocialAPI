package api

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/mediatR"
	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/melody"
	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/service"
	"github.com/google/uuid"
	mel "github.com/olahol/melody"

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

type gopherInfo struct {
	ID, X, Y string
}

func StartHttpServer(lc fx.Lifecycle, db *domain.Queries, logger *zap.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	v1 := apiVersioning(e)
	AddController(v1, db)

	// Serve websocket demo html application
	wd, _ := os.Getwd()
	e.File("/demo", filepath.Join(wd, "/static/index.html"))

	// Handle websocket
	var websocket = melody.NewWebsocket()

	e.GET("/ws", func(c echo.Context) error {
		return websocket.HandleRequest(c.Response(), c.Request())
	})

	websocket.HandleConnect(func(s *mel.Session) {
		ss, _ := websocket.Sessions()

		for _, o := range ss {
			value, exists := o.Get("info")

			if !exists {
				continue
			}

			info := value.(*gopherInfo)

			s.Write([]byte("set " + info.ID + " " + info.X + " " + info.Y))
		}

		id := uuid.NewString()
		s.Set("info", &gopherInfo{id, "0", "0"})

		s.Write([]byte("iam " + id))
	})

	websocket.HandleDisconnect(func(s *mel.Session) {
		value, exists := s.Get("info")

		if !exists {
			return
		}

		info := value.(*gopherInfo)

		websocket.BroadcastOthers([]byte("dis "+info.ID), s)
	})

	websocket.HandleMessage(func(s *mel.Session, msg []byte) {
		p := strings.Split(string(msg), " ")
		value, exists := s.Get("info")

		if len(p) != 2 || !exists {
			return
		}

		info := value.(*gopherInfo)
		info.X = p[0]
		info.Y = p[1]

		websocket.BroadcastOthers([]byte("set "+info.ID+" "+info.X+" "+info.Y), s)
	})

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
