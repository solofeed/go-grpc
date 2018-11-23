start:
	docker-compose up csv-service
	make show-users

show-users:
	docker-compose exec mongodb mongo test --eval "db.users.find().limit(5).forEach(printjson)"


.PHONY: start show-users
