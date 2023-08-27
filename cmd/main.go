package main

import (
	"github.com/anacondaf/petSocialAPI/src/api/host/domain"
	"github.com/anacondaf/petSocialAPI/src/api/infrastructure/logger"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"os"
	"path/filepath"

	api "github.com/anacondaf/petSocialAPI/src/api/host"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func connectElasticSearch() {
	wd, _ := os.Getwd()
	cert, err := os.ReadFile(filepath.Join(wd, "/elastic/certs/ca/ca.crt"))
	if err != nil {
		log.Fatal(err)
	}
	config := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "Password",
		CACert:   cert,
	}

	es, err := elasticsearch.NewClient(config)
	if err != nil {
		return
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res.Body)
}

func main() {
	fx.New(
		fx.Provide(
			domain.UseSqlc,
			logger.UseZap,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(
			api.StartHttpServer,
			connectElasticSearch),
	).Run()
}
