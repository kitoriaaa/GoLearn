FROM golang:1.15.3
WORKDIR /go/src/app
RUN apt-get update
RUN apt-get install -y sqlite3
RUN go get github.com/go-sql-driver/mysql