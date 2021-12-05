build:
	docker build --no-cache --rm -t yvv4docker/goecho:v1.0.0 -f Dockerfile .
	docker system prune -f

run:
	docker run --rm -p8080:8080 --name goecho yvv4docker/goecho:v1.0.0

stop:
	docker stop goecho

push:
	docker push yvv4docker/goecho:v1.0.0

check:
	curl -X GET 'http://localhost:8080/v1/address'