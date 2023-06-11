pre_install:
	go mod download
	docker-compose up -d

worker:
	go run worker/main.go

dash:
	go run dashboard/main.go

app:
	go run app/server.go


