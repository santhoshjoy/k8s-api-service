FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o api-server main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api-server .
EXPOSE 8080
CMD ["./api-server"]
