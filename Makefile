build:
	docker build --no-cache --rm -t yvv4docker/goecho:v1.0.0 -f Dockerfile .

run:
	docker run --rm -p8080:8080 --name goecho yvv4docker/goecho:v1.0.0

push:
	docker push yvv4docker/goecho:v1.0.0

check:
	curl -X GET 'http://localhost:8080/v1/address'