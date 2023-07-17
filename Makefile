dev:
	air

migrateup:
	migrate -path app/database/migrations -database ${database} up

migratedown: 
	migrate -path app/database/migrations -database ${database} down

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: migrateup migratedown proto