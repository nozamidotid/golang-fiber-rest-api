FROM golang:alpine3.19

WORKDIR /app
COPY . .
RUN go mod tidy