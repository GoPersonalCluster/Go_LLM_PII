build:
	docker-compose up -d
build-reuse:
	docker compose up -d postgres-pii-llm llm app-llm-pii
bash-app:
	docker exec -it app-llm-pii bash

reset-docker:
	@docker stop $$(docker ps -aq) 2>/dev/null || true
	@docker rm $$(docker ps -aq) 2>/dev/null || true
	@docker rmi $$(docker images -aq) 2>/dev/null || true
	@docker volume rm $$(docker volume ls -q) 2>/dev/null || true
	@docker network prune -f