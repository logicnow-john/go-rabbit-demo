FROM golang:1.17-alpine3.14 AS builder
WORKDIR /go/src/app

COPY ../../go.mod go.sum ./
RUN go mod download

COPY ../.. .
RUN go build -o /go/bin/server ./cmd/server/

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /go/bin/server .
COPY --from=builder /go/src/app/cmd/server/app.env /cmd/server/app.env
USER 65534

ENTRYPOINT [ "./server" ]