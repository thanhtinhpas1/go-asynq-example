# 📖 Example: Asynq. Simple, reliable & efficient distributed task queue for your next Go project

[Asynq](https://github.com/hibiken/asynq) is a Go library for distributed tasks and processing them asynchronously with multiple workers. It's backed by Redis and is designed to scale and distribute easily.
Asynq has many features for task like schedule, timeout, retry, etc.

![Overview](https://user-images.githubusercontent.com/11155743/116358505-656f5f80-a806-11eb-9c16-94e49dab0f99.jpg)

## Prerequisites
Make sure redis already exists in your environment or install by using below commands:

```console
brew install redis
```
or
```console
docker-compose up -d
```

## Quick start
1. Clone this repository
```console
git clone github.com/thanhtinhpas1/go-asynq-example.git
```
2. Install dependencies
```console
go mod download
```
or Makefile
```console
make pre_install
```
3. Start worker
```console
go run worker/main.go
```
or Makefile
```console
make worker
```
4. Start dashboard UI
```console
go run dashboard/main.go
```
or Makefile
```console
make dash
```

5. Start app-server
```console
go run app/server.go
```
or Makefile
```console
make app
```

If you like my example, please give me a star ⭐
