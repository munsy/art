# Golang build container
FROM golang:1.13.5

WORKDIR $GOPATH/src/github.com/munsy/art

RUN go build

COPY art art

ENTRYPOINT [ "./art" ]
