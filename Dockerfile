FROM golang:alpine

ADD /src /

WORKDIR /

RUN go build -o hermes .

CMD ["/hermes"]
