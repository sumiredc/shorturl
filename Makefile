include .env

.PHONY: swagger
swagger: ## swagger
	@npx swagger-typescript-api -p ./openapi/index.yaml -o ./front/src/lib -n api.ts

.PHONY: mysql
mysql: ## MySQLにログインする
	@docker compose exec mysql mysql -h $(MYSQL_HOST) -u $(MYSQL_USER) -p$(MYSQL_PASSWORD) $(MYSQL_DATABASE)

.PHONY: migrate
migrate: 
	@migrate --path database/migrations --database 'mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' -verbose up

.PHONY: rollback
rollback: 
	@migrate --path database/migrations --database 'mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' down
