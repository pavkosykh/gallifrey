build:
	go mod download && go build -o ./.bin/gallifrey-cli ./cmd/cli/main.go