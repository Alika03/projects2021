Run the following commands:

1. docker image build -t ascii-art-web-docker .

2. docker images

3. docker container run -d --name=ascii-art-web -p 8070:8070 ascii-art-web-docker

4. docker exec -it ascii-art-web ls -l

5. docker system prune