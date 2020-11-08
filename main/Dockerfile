FROM golang:latest

ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go

RUN mkdir -p $GOPATH/src/support/
WORKDIR $GOPATH/src/support/

COPY . $GOPATH/src/support/

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

EXPOSE 8080

RUN go build -o main .

RUN ls

CMD ["go", "run", ".", "&"]
