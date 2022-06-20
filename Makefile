compose-restart:
	docker-compose down
	docker rmi $(docker images -a -q)
	docker-compose up --force-recreate