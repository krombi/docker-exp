FROM golang:latest

ENV CONTAINER_PATH = /usr/src/support/

RUN mkdir -p ${CONTAINER_PATH}
WORKDIR ${CONTAINER_PATH}

COPY . ${CONTAINER_PATH}

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

EXPOSE 8080

CMD ["go", "run", "main.go", "&"]
