# web-go-gin



## Run

```bash
go run ./cmd/server        # http://localhost:8000/health
```

`PORT` overrides the listen port (default `8000`) — the platform sets it when it runs the app.

## Test / lint

```bash
go test ./...
go vet ./...
gofmt -l .
```

## Layout

```
cmd/server/main.go        main only: env, app.New, Run
internal/
├── app/app.go            router and every route; Version
├── handlers/             HTTP: parse the request, call a service, write the response
│   ├── health.go         GET /health
│   └── items.go          the example resource: list, create, get
├── services/             business rules; never import gin; return errors handlers translate
└── models/               the structs that cross the HTTP boundary
```

Request → route → handler → service → response. A handler holds no rules; a service holds no HTTP. Add `repositories/` when a database or an external API appears, and `middleware/` when a route needs auth or logging.
