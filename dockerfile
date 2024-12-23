#build stage
FROM golang:1.23-alpine3.19 AS builder
RUN apk add --no-cache git upx
WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download
COPY . .
RUN go build -o app .
RUN upx app

#final stage
FROM alpine:latest
RUN apk update && apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app .
ENTRYPOINT ["./app"]