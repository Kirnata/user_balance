build:
	docker build -t app .
run:
	docker run -d --name app -p 80:8081 app
stop:
	docker stop app

