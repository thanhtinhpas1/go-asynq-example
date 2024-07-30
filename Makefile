.PHONY: pre_install worker dash app

pre_install:
	go mod download

worker:
	go run worker/main.go

dash:
	go run dashboard/main.go

app:
	go run app/server.go


