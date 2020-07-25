FROM golang:alpine

ADD server.go /

WORKDIR /

RUN go get github.com/paulmj7/hermes/hermes

RUN go build -o hermes .

CMD ["/hermes"]
