GOPATH:=$(shell go env GOPATH)

.PHONY: run build-consignment

build-consignment:
	cd consignment && make build && rm consignment-srv

build-vessel:
	cd vessel && make build && rm vessel-srv

run: build-consignment build-vessel
	docker-compose up -d consignment vessel microweb

seed-db-local:
	@docker cp  data/seed.sql  lotus_db:/seed.sql
	@docker exec -t lotus_db sh -c "PGPASSWORD=lotus psql -U lotus -d lotus -f /seed.sql"

init: remove-infras run seed-db-local

remove-infras:
	docker-compose stop; docker-compose rm -f