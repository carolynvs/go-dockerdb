FROM golang:1.10-stretch

RUN mkdir -p /go/src/github.com/carolynvs/go-dockerdb
WORKDIR /go/src/github.com/carolynvs/go-dockerdb

COPY . .

CMD ["go", "run", "main.go"]