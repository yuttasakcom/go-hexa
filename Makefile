run:
	go run main.go

dev:
	air

migrateup:
	migrate -path app/database/migrations -database ${database} up

migratedown: 
	migrate -path app/database/migrations -database ${database} down

proto:
	protoc --proto_path=src/domain --go_out=src/domain --go_opt=paths=source_relative \
	--go-grpc_out=src/domain --go-grpc_opt=paths=source_relative \
	src/domain/**/*.proto

.PHONY: migrateup migratedown proto