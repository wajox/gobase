# image for compiling binary
ARG BUILDER_IMAGE="golang:1.16-alpine"
# here we'll run binary app
ARG RUNNER_IMAGE="alpine:3.12"

FROM ${BUILDER_IMAGE} as builder
### variables
# disable golang package proxying for this modules. necessary for private repos. Example: github.com/company-name
ARG GOPRIVATE=""
# ssh key for accessing private repos
ARG SSH_KEY

ENV GO111MODULE on
ENV GOPRIVATE ${GOPRIVATE}

RUN apk add --update gcc g++ openssh git make

# configure git to work with private repos
RUN mkdir -p ~/.ssh/ &&\
    echo "${SSH_KEY}" >> ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa && ssh-keyscan github.com >> ~/.ssh/known_hosts &&\
    git config --global url."git@github.com:".insteadOf "https://github.com/"

### copying project files
WORKDIR /go/src/github.com/wajox/gobase
# copy gomod 
COPY go.mod go.sum ./
# Get dependancies. Also will be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# creates build/main files
RUN make build

FROM ${RUNNER_IMAGE}

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
