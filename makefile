setup:
	docker-compose up -d
setup-down:
	docker ps -a --format "{{.ID}} {{.Names}}" | grep mysql-crud-api | awk '{print $$1}'| xargs docker stop | xargs docker rm -v
run:
	go run main.go