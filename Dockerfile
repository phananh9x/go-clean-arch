FROM golang:1.17-alpine as builder
ARG GITHUB_TOKEN

ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add --update --no-cache ca-certificates curl git
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH/src/go-clean-arch

COPY . .
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod download  && go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /main ./cmd/service

FROM alpine:3.7

RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /main /app/main

COPY build/app.yaml /app/

CMD ["./main", "--config-file", "./app.yaml"]

