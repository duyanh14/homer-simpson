FROM golang:1.18.2-alpine as builder
WORKDIR /app
COPY . /app

ENV BUILD_TAG 1.0.0
ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go mod vendor
RUN go build -o simpson /app/cmd/server/main.go
FROM golang:1.18.2-alpine
WORKDIR /app
COPY --from=builder /app/simpson /app/simpson.go
CMD ["./simpson.go"]
