FROM golang:1.23 AS builder
ENV CGO_ENABLED=0

WORKDIR /go

COPY go.mod .
COPY go.sum .
COPY .env .
COPY . .
RUN go mod download

RUN go build -o bysykkel main.go

FROM gcr.io/distroless/base:nonroot
WORKDIR /go
USER nonroot

COPY --from=builder /go/bysykkel bysykkel

ENTRYPOINT ["/go/bysykkel"]
