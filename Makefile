build-docker-image:
	docker build -t madebyais/bucket .

run-docker-local:
	docker run -p 8700:8700 -v /opt/bucket:/opt/bucket -v /etc/bucket:/etc/bucket madebyais/bucket