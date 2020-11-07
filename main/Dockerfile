FROM golang:latest

ENV GOPATH $HOME/go

RUN mkdir -p $HOME/go/src/support/
WORKDIR $HOME/go/src/support/

COPY . $HOME/go/src/support/

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

EXPOSE 8080

RUN go build

CMD ["go run", "&"]
