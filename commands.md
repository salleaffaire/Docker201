# List of commands

## Writing a dockerized Go application

countw is a simple Go application that counts the number of words in a file. It reads the file from the file system and prints the number of words to the console.

First we build it on the host,

```go
go build -o countw countw.go
```

Then we run it,

```bash
./countw test.txt
```

We then containerize it using Docker. We create a Dockerfile with the following content,

```Dockerfile
FROM golang:1.18.8-alpine3.16

LABEL version=1.0.0

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o countw

# ENTRYPOINT ["./countw"]
CMD ["sh"]
```

We pull the base image first,

```bash
docker pull golang:1.18.8-alpine3.16
```

We build the image using the following command,

```bash
docker build -t demo-go-app .
```

We can then run the container using the following command,

```bash
docker run --rm -it --name demo-go-app demo-go-app
```

Once we do that we realize that we need to share the file with the container. We can do that using volumes.

We launch the container using the following command, mounting the current directory to the /app/mount directory in the container,

```bash
docker run --rm -it --name demo-go-app -v `pwd`:/app/mount demo-go-app
```

Then update the ENTRYPOINT in the Dockerfile to the following,

And rebuild the image,

```bash
docker build -t demo-go-app .
```

We can then run the container using the following command,

```bash
docker run --rm -it --name demo-go-app -v `pwd`:/app/mount demo-go-app mount/test.txt
```

We can also use multi-stage builds to reduce the size of the image.

```bash
docker pull golang:1.18.8-alpine3.16

docker build -t demo-go-app .

docker run --rm -it --name demo-go-app demo-go-app

docker run --rm -it --name demo-go-app -v `pwd`:/app/mount demo-go-app mount/test3.txt

# Volumes

docker volume create my-volume
docker volume ls
docker volume inspect my-volume
docker run --rm -it --name demo-go-app -v my-volume:/app/mount demo-go-app
docker volume rm my-volume

docker history demo-go-app

# Multi-stage build

# Networks

# On the host 
ip address show # Point to the bridge network called "docker0"

docker run --rm -dit --name superman busybox sh
docker run --rm -dit --name hulk busybox sh
docker run --rm -dit --name batman nginx sh

ip address show
bridge link show

docker exec -it superman sh

# Once in the container
ip address show

ping 172.18.0.3 # Or whatever the IP of the batman container is
ping google.com

ip route show # Shows the default gateway

# Now host network

# Stop batman
docker container stop batman

# start batman on bridge with port exposed
docker run --rm -dit -p 8080:80 --name batman nginx

# Back on the h
curl localhost:8080

# or 
curl http://172.18.0.4:80 // The IP of the batman container


# Start batman with host network
docker run --rm -dit --network host --name batman nginx

# Back on the host
# Now create a new network
docker network create --driver bridge pluto

docker network ls

ip address show

docker run --rm -dit --network pluto --name earth busybox sh
docker run --rm -dit --network pluto --name mars busybox sh

docker network ls

docker run --rm -dit --name alpine1 alpine ash
docker run --rm -dit --name alpine2 alpine ash

docker network create --driver bridge alpine-net

docker network ls

docker network inspect alpine-net

docker run --rm -dit --name alpine1 --network alpine-net alpine ash
docker run --rm -dit --name alpine2 --network alpine-net alpine ash
docker run --rm -dit --name alpine3 alpine ash
docker run --rm -dit --name alpine4 --network alpine-net alpine ash
docker network connect bridge alpine4

docker network inspect alpine-net
docker network inspect bridge

docker container attach alpine4

# inside
ip addr ls

ping -c 2 alpine2
ping -c 2 alpine4
ping -c 2 alpine1 # Itself
pint -c 2 alpine3 # Won't work

# alpine3 not visible at all through alpine1

docker container stop alpine1 alpine2 alpine3 alpine4

```
