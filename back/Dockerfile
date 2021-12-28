FROM golang:1.16

RUN mkdir /back
WORKDIR /back

RUN go mod init github.com/obrkn/twitter
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/cosmtrek/air