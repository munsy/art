FROM golang:alpine

RUN mkdir /app 

ADD . /app/

WORKDIR /app 

RUN go build -o art .
RUN adduser -S -D -H -h /app art

USER art

EXPOSE 8080

CMD ["./art", "serve", "-p", "8080"]