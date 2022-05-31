FROM golang:1.18-alpine

WORKDIR /app

COPY app .

EXPOSE 8080

CMD [ "go", "run", "main.go" ]