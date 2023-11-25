FROM golang:1.21 AS base
RUN apt update && apt install -y git make ca-certificates 

# Code generators
ENV CGO_ENABLED=1
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.24.0
RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app
RUN apt update && apt-get install -y postgresql-client vim wget openssh-server
#RUN mkdir /var/run/sshd
#RUN echo 'root:password' | chpasswd
#RUN sed -i 's/PermitRootLogin without-password/PermitRootLogin yes/' /etc/ssh/sshd_config

# SSH login fix. Otherwise user is kicked off after login
#RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

# Redis CLI from source
RUN wget https://download.redis.io/redis-stable.tar.gz && tar -xzf redis-stable.tar.gz && cd redis-stable && make install
RUN rm -fr redis-stable redis-stable.tar.gz

# Dev tools
#RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Hot reloading tools
RUN go install github.com/cosmtrek/air@latest 
EXPOSE 22
EXPOSE 8080
EXPOSE 9000

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY go.mod /app/
COPY go.sum /app/
RUN go mod download \
    && go mod verify

COPY . /app
RUN go build -o bin/comics -a ./cmd/server/comics/

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/bin/comics /usr/local/bin/comics
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/comics"]
