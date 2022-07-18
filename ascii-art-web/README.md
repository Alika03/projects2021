# ASCII-ART-WEB

## Author: Alika96

## How to run
Run the following commands:
1. For building an image: **docker image build -t ascii-art-web-docker .**
2. For showing images: **docker images**
3. For running a container: **docker container run -d --name=ascii-art-web -p 8070:8070 ascii-art-web-docker**
4. For showing files inside the running container: **docker exec -it ascii-art-web ls -l**


# Image
1. ![alt text](https://github.com/Alika03/ascii-art-web/blob/a06633682b9e617a42b5ec05f953b9d7f74e793b/imageofweb/main.png?raw=true)

Then:
1. Open **localhost:8070** in any browser.
2. Choose any banner you want.
3. Type some text and click "Нажать" for printing the text in ASCII-art or click "Скачать" for saving the art as a file in .txt format 

## Instructions

* The web server is created in **Go.**
* It uses Docker.
* The project has a **Dockerfile.**
* Your Dockerfile respects the [Dockerfile good practices.](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

## Usage
* You can see all about Docker on [docker docs](https://docs.docker.com/).

This project contains:

* Client utilities.
* The basics of web :
    * Server
    * HTML
    * HTTP
* Ways to receive data.
* Ways to output data.
* [docker](https://docs.docker.com/).
* Using and [setting up Docker](https://docs.docker.com/get-started/) :
    * Services and dependencies.
    * Containerizing an application.
    * Compatibility/Dependency.
    * Creating images.

