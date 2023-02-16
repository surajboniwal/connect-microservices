build_gateway:
	cd gateway && make build && cd ..

remove_gateway:
	cd gateway && make remove && cd ..

build_authentication:
	cd authentication && make build && cd ..

remove_authentication:
	cd authentication && make remove && cd ..

build: build_gateway build_authentication

remove: remove_gateway remove_authentication

up: remove build
	docker-compose up --detach --build

down: remove
	docker-compose down --rmi local