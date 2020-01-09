FROM golang:alpine AS builder
MAINTAINER docker@munsy.io
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o art .

FROM node:alpine
WORKDIR /app
COPY --from=builder /app/ .
RUN npm i
COPY . .
EXPOSE 80 5000
CMD ./art serve -p $PORT 