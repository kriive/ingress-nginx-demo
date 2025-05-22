FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY ./go.mod .
COPY ./main.go .

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=builder /go/bin/app /
CMD ["/app"]
