FROM golang:latest

RUN mkdir -p /app/support/
WORKDIR /app/support/

COPY . /app/support/

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket

EXPOSE 8080

RUN go build -o main .
CMD ["./main", "&"]
