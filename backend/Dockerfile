FROM golang:1.16-alpine AS builder
WORKDIR /
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/ .

FROM alpine:3
COPY --from=builder docs/swagger.yaml ./docs/
COPY --from=builder /out/ .
ENTRYPOINT ["./Open-Stage"]
