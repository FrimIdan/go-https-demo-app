FROM golang:1.16.6-alpine AS builder

#RUN apk add --update --no-cache gcc g++ make git openssh
RUN apk add --update --no-cache gcc g++ make

WORKDIR /build
COPY go.* ./
RUN go mod download

# Copy and build https-test code
COPY . .
RUN go build -o https-test ./cmd/https-test/main.go


FROM alpine:3.14.0

WORKDIR /app

COPY --from=builder ["/build/https-test", "./https-test"]

ENTRYPOINT ["/app/https-test"]
