FROM node:alpine AS builder
WORKDIR /app
COPY ../../ .
RUN npm install && npm run docker
COPY . .

FROM nginx:alpine
COPY --from=builder /app/dist/site/ /usr/share/nginx/html/
COPY --from=builder /app/dist/site/assets/jfrog-logo.png /usr/share/nginx/html/art/assets/jfrog-logo.png
EXPOSE 80