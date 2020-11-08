FROM golang:latest

RUN mkdir -p /app/support/
WORKDIR /app/support/

COPY . /app/support/

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/websocket
RUN go get -u github.com/go-sql-driver/mysql

EXPOSE 8080

RUN go build -o main .
CMD ["./main", "&"]
