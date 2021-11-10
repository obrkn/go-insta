FROM golang:1.16

RUN mkdir /back
WORKDIR /back

RUN go mod init github.com/obrkn/go-insta
RUN go get -u github.com/go-sql-driver/mysql

CMD go run cmd/go-insta/main.go