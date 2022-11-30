build-backend:
	sudo docker image build -t muhammadrivaldy05/dealljobs .

build-swagger:
	sudo docker image build -t muhammadrivaldy05/swagger-ui -f Dockerfile.swagger-ui .

push-image-backend:
	sudo docker image push muhammadrivaldy05/dealljobs

push-image-swagger:
	sudo docker image push muhammadrivaldy05/swagger-ui

remove-useless:
	sudo docker image prune

migration-security:
	@read -p "input the migration name: " name; \
	migrate create -ext sql -dir handler/security/entity/database/migrations $$name

migration-user:
	@read -p "input the migration name: " name; \
	migrate create -ext sql -dir handler/user/entity/database/migrations $$name