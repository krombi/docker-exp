FROM golang:latest

ENV GOPATH /go

RUN mkdir -p /go/src/support/
WORKDIR /go/src/support/

COPY . /go/src/support/

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

EXPOSE 8080

RUN go build

CMD ["go run", "&"]
