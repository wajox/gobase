# Start from a Debian image with Go 1.15 installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.16-alpine as builder

RUN apk add --update gcc g++ openssh git make

ARG SSH_KEY

# setup dependencies
ADD . /go/src/github.com/wajox/gobase
WORKDIR /go/src/github.com/wajox/gobase

ENV GO111MODULE on

# creates build/main files
RUN mkdir -p ~/.ssh/ &&\
    echo "${SSH_KEY}" >> ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa && ssh-keyscan github.com >> ~/.ssh/known_hosts &&\
    git config --global url."git@github.com:".insteadOf "https://github.com/" &&\
    make build

FROM alpine:3.12

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories &&\
    apk update &&\
    apk add --no-cache\
    ca-certificates

RUN mkdir -p ./api
RUN mkdir -p ./db/migrations

COPY --from=builder /go/src/github.com/wajox/gobase/api ./api
COPY --from=builder /go/src/github.com/wajox/gobase/db/migrations ./db/migrations

COPY --from=builder /go/src/github.com/wajox/gobase/build/app .

CMD ["./app", "s"]
