FROM golang:alpine 
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o art .
EXPOSE 5000
CMD ["./art"]