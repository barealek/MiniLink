# I'm using doppler to manage my secrets, you can use your own method or just use the .env file.
# Adjust this script to your needs.

up:
	doppler run -- docker compose up -d --build --remove-orphans
build:
	docker compose build
down:
	docker compose down
kill:
	docker compose kill
restart:
	docker compose kill
	doppler run -- docker compose up -d
logs:
	docker compose logs -f
update:
	git pull
	docker compose build
	make restart