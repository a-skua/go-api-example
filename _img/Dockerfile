FROM golang:1.18 AS builder

ADD ./src /go/src

WORKDIR /go/src

ARG CGO_ENABLED=0
RUN go build -o /go/bin/api ./cmd/api.go

FROM scratch

COPY --from=builder /go/bin/api /bin/api

ENTRYPOINT ["/bin/api"]
