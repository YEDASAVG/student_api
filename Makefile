include .env
export

run:
	go run cmd/main.go
build: 
	go build -o student_api cmd/main.go
test:
	go test ./...	
migrate-up:
	migrate -path migrations -database ${DATABASE_URL} -verbose up
migrate-down:
	migrate -path migrations -database ${DATABASE_URL} -verbose down 1