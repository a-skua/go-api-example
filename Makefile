up:
	docker compose up -d api

build:
	docker compose build api

coverage: test
	docker compose run --rm gopher bash -c 'cd src; go tool cover -html=/go/cover/c.out -o=/go/cover/index.html'

test:
	docker compose run --rm gopher bash -c 'cd src; go test -cover -coverprofile=/go/cover/c.out ./...'

fmt:
	docker compose run --rm gopher bash -c 'cd src; go fmt ./...'

tidy:
	docker compose run --rm gopher bash -c 'cd src; go mod tidy'

mysql:
	docker compose exec db mysql -uroot -p api_example

mysql-log:
	docker compose exec db tail -f /var/lib/mysql/general.log

migrate:
	docker compose run --rm migrate

rollback:
	docker compose run --rm migrate db:rollback

migrate-tasks:
	docker compose run --rm migrate -T

e2e:
	_e2e/test.sh
