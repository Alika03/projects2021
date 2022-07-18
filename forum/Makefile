build:
	go build -o bin/main main.go
run:
	go run main.go
docker:
	docker image build -f Dockerfile -t goo .
	docker container run -p 9090:9090 -d --name forum goo
