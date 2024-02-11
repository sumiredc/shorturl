# README

GO x Next.js

## Setup

```shell
# GO(backend)
cd back
go mod tidy

# Next.js(frontend)
cd front
pnpm install

# Open API
make swagger # front/src/lib/api.ts の作成
```

## Flow

### How to create api
1. GO で APIを作成
2. OPEN API で エンドポイントを定義
3. Next.js からAPIをコール

## Next.js

- https://nextjs.org/
- https://mui.com/

```shell
cd front
pnpm dev
```

## Go

- https://go.dev/
- https://echo.labstack.com/
- https://gorm.io/index.html
- https://github.com/cosmtrek/air

```shell
cd back

go run main.go
# or
air
```

## Migration

- https://github.com/golang-migrate/migrate

### Create

```shell
migrate create -ext sql -dir database/migrations -seq create_***_table
```

### Run

```shell
migrate --path database/migrations --database 'mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' -verbose up
```

### Rollback

```shell
migrate --path database/migrations --database 'mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' down
```
