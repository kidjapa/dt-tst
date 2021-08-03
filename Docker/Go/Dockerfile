# Start from golang base image
FROM golang:1.16.5 AS builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /src

WORKDIR /src

RUN pwd
# First add modules list to better utilize caching
COPY go.sum go.mod /src/

# Download dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build components.
# Put built binaries and runtime resources in /app dir ready to be copied over or used.
RUN go install -installsuffix cgo -ldflags="-w -s" && \
    mkdir -p /app && \
    cp -r $GOPATH/bin/dt-tst /src/

#
# 2. Runtime Container
#
FROM alpine:latest

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

RUN apk add --update --no-cache \
    tzdata \
    ca-certificates \
    bash \
    && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

COPY --from=builder /src /app/

CMD ["./dt-tst"]