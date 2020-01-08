FROM golang:alpine AS builder
MAINTAINER docker@munsy.io
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o art .

FROM node:alpine
COPY --from=builder /app/ /app/
WORKDIR /app
COPY . .
RUN npm install && npm run docker
COPY . .
EXPOSE 5000
CMD ["./art", "serve", "--angular", "--container"]