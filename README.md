## Migration

### Create

```shell
migrate create -ext sql -dir database/migrations -seq create_***_table
```

### Run

```shell
migrate --path database/migrations --database 'mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' -verbose up
```


