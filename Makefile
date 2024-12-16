migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' down

show_migrate:
	psql -h 0.0.0.0 -p 5436 -U postgres -d postgres

build:
	docker-compose up --build

run:
	docker-compose up

down:
	docker-compose down

test:
	go test -v ./...