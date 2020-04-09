# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM djirik/gokafka:1.14-alpine as builder
RUN apk add gcc g++ make

ARG SSH_KEY

# setup dependencies
ADD . /go/src/github.com/ildarusmanov/gobase
WORKDIR /go/src/github.com/ildarusmanov/gobase

ENV GO111MODULE on

# creates build/main files
RUN mkdir -p ~/.ssh/ &&\
    echo "${SSH_KEY}" >> ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa && ssh-keyscan github.com >> ~/.ssh/known_hosts &&\
    git config --global url."git@github.com:".insteadOf "https://github.com/" &&\
    make build

FROM djirik/alkafka

RUN mkdir -p ./docs
RUN mkdir -p ./db/migrations
RUN mkdir -p ./app/config

COPY --from=builder /go/src/github.com/ildarusmanov/gobase/docs ./docs
COPY --from=builder /go/src/github.com/ildarusmanov/gobase/db/migrations ./db/migrations

COPY --from=builder /go/src/github.com/ildarusmanov/gobase/build/main .


CMD ["./main", "s"]
