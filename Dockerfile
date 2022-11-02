FROM golang:1.19.2-alpine3.16
WORKDIR /app
COPY . .
RUN go build -o main

EXPOSE 8081
CMD ["/app/main"]