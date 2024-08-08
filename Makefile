build:
	go build -o ./bin/game ./cmd/main.go

run:build
	./bin/game
