run-dev-env:
	docker compose -f ./docker-compose.dev.yml --env-file .env up --build -d

test:
	go test ./...
