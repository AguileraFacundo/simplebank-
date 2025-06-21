createdb:
	docker exec -it -u 0 173d0f825da5 psql -c "CREATE DATABASE simple_bank;"

dropdb:
	docker exec -it -u 0 173d0f825da5 psql -c "DROP DATABASE simple_bank;"
	
databases:
	docker exec -it -u 0 173d0f825da5 psql -c "\l"

dev:
	air

.PHONY: createdb dropdb databases dev