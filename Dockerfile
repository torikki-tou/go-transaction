FROM golang:1.18

WORKDIR /usr/src/app

COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify

COPY src .
RUN go build -v -o /usr/local/bin/app main.go

CMD ["app"]