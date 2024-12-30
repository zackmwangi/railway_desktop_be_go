# Welcome to the Railway desktop app API in golang

This API exposes RESTful/gRPC endpoints to be consumed by a mobile/web frontend.

The API utilizes gRPC gateway for dual-accessibility.

The frontend interface is documented in swagger OpenAPI spec.

This API is an intermediary service converting requests into GraphQL queries sent to Railway.app APIs

```

go mod tidy

go run cmd/main.go

```

