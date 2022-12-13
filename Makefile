build:
	docker build -t app .
run:
	docker run --name app -p 80:8080 app
stop:
	docker stop app

