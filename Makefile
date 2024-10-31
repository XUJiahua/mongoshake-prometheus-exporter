docker_build:
	docker build --platform linux/amd64 -t johnxu1989/mongoshake-prometheus-exporter:0.1.0 .

docker_push:
	docker push johnxu1989/mongoshake-prometheus-exporter:0.1.0
