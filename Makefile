dev:
	air

build:
	go build -o ./bin/main ./cmd/main.go

run:
	go run cmd/main.go

sqlc:
	docker run --rm -v /c/Users/adisu/OneDrive/Documents/NERD/PROJECTS/Perpustakaan-LenU:/src -w /src kjconroy/sqlc generate

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:secret@localhost:5433/perpustakaan_lenu?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:yungbenn@localhost:5432/perpustakaan_lenu?sslmode=disable" down