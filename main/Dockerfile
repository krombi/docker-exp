FROM golang:latest

RUN mkdir -p /usr/src/support/
WORKDIR /usr/src/support/

COPY . /usr/src/support/

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

EXPOSE 8080

RUN go run main.go &

# CMD ["./main", "&"]
