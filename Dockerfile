FROM golang:1.18.8-alpine3.16

LABEL version=1.0.0

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o countw

# ENTRYPOINT ["./countw"]
CMD ["sh"]