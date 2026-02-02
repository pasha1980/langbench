FROM golang:1.25.5 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/main -o main

FROM scratch as runner
COPY --from=builder /app/main /main
CMD ["/main"]
