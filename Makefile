DB_SOURCE=mysql://root:P@ssword!@tcp(localhost:3306)/pet-social
MIGRATE_DIR=E:\Projects\Go_Projects\go-echo\api\migrations

sqlc-gen:
	cd ./cmd/internal/sqlc && sqlc generate

migrate:
	docker run -v "$(MIGRATE_DIR)":/migrations --network host migrate/migrate -path=/migrations/ -database "$(DB_SOURCE)" up

gen-migrate:
	migrate create -ext sql -dir "$(MIGRATE_DIR)" -seq $(NAME)

.PHONY: sqlc-gen migrate gen-migrate