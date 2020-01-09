FROM golang:alpine AS builder
MAINTAINER docker@munsy.io
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o art .

FROM node:alpine
WORKDIR /app
COPY --from=builder /app/ .
RUN npm i && $(pwd)/node_modules/.bin/ng build --configuration=docker
EXPOSE 80 5000
RUN ls -l
CMD ./art serve -p $PORT 