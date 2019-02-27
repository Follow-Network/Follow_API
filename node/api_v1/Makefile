docker_name := followinc_api_v1_docker
docker_tag := 1.0
container_name := followInc_api_v1

port := 5000

docker-no-cache:
	docker build --no-cache -t ${docker_name}:${docker_tag} -f Dockerfile ./

docker:
	docker build -t ${docker_name}:${docker_tag} -f Dockerfile ./

run:
	docker run -p ${port}:5000 --rm -d -it --name ${container_name} ${docker_name}:${docker_tag}

run-no-d:
	docker run -p ${port}:5000 --rm -it --name ${container_name} ${docker_name}:${docker_tag}

Nolan:
	docker exec -it ${container_name} bash

stop:
	docker stop ${container_name}

logs:
	docker logs ${container_name}


delete-container:
	docker images
	docker rmi ${docker_name}:${docker_tag}
	docker images
