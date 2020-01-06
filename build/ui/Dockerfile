FROM node:alpine AS builder
WORKDIR /app
COPY . .
RUN npm install && npm run build
COPY . .

FROM nginx:alpine
COPY --from=builder /app/dist/site/* /usr/share/nginx/html/art/
EXPOSE 80