run:
	go run cmd/main.go
build: 
	go build -o student_api cmd/main.go
test:
	go test ./...	
