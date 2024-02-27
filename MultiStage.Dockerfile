FROM golang:1.18.8-alpine3.16 AS BuildStage

LABEL version=1.0.0

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o countw

FROM alpine:latest

WORKDIR /app

COPY --from=BuildStage /app/countw /app/countw

CMD ["sh"]
