FROM golang:latest

ENV USR_PATH = /usr/
ENV GOPATH = ${USR_PATH}

ENV CONTAINER_PATH = ${USR_PATH}src/support/

RUN mkdir -p ${CONTAINER_PATH}
WORKDIR ${CONTAINER_PATH}

COPY . ${CONTAINER_PATH}

RUN go get -u github.com/gorilla/mux \
    go get -u github.com/gorilla/websocket

EXPOSE 8080

CMD ["go", "run", "main.go", "&"]
