build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.account . 

run:
	docker service create --name account-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6014:6014 --publish 7014:7014 github.com/turistikrota/service.account:latest

remove:
	docker service rm account-api-turistikrota-com

stop:
	docker service scale account-api-turistikrota-com=0

start:
	docker service scale account-api-turistikrota-com=1

restart: remove build run
