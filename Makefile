DB_SOURCE=mysql://root:P@ssword!@tcp(localhost:3306)/pet-social
MIGRATE_DIR=E:\Projects\Go_Projects\go-echo\api\migrations
ELASTIC_PASSWORD=Password

run:
	nodemon --exec "go run cmd/main.go" --signal SIGTERM

sqlc-gen:
	cd ./cmd/internal/sqlc && sqlc generate

migrate:
	docker run -v "$(MIGRATE_DIR)":/migrations --network host migrate/migrate -path=/migrations/ -database "$(DB_SOURCE)" up

gen-migrate:
	migrate create -ext sql -dir "$(MIGRATE_DIR)" -seq $(NAME)

test-es1:
	curl --cacert elastic/certs/ca/ca.crt -u elastic:Password https://localhost:9200

.PHONY: run sqlc-gen migrate gen-migrate test-es1