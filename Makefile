build:
	go build -o bin/app cmd/app/main.go

run:
	cd cmd/app && go run main.go