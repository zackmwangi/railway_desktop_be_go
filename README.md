# Welcome to the Railway desktop app API in golang

This API exposes RESTful/gRPC endpoints to be consumed by a personalized mobile/web frontend.

The API is gRPC first and utilizes a gRPC gateway for dual-accessibility, mapping RESTful requests onto gRPC logic.

The frontend interface is documented in swagger OpenAPI spec.

This API is an intermediary service converting requests into GraphQL queries that are sent on to Railway.app APIs

```

go mod tidy

go run cmd/main.go

```

