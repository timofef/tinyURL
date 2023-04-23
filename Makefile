compose-in-memory:
	sudo IN_MEM=-in-memory docker compose up -d --no-deps tinyurl

compose-postgres:
	sudo docker compose up -d
