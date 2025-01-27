FROM golang:1.23.1

RUN apt-get update && apt-get install -y tzdata

ENV TZ=Europe/Paris

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go mod download

ENTRYPOINT CompileDaemon --build="go build ./cmd/server/main.go" --command=./main