FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary

FROM scratch
COPY --from=builder /app-binary /go/bin/app
ENV HTTP_PORT=80
ENTRYPOINT ["/go/bin/app"]
