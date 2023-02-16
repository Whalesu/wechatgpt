FROM golang:1.19-alpine as builder

ARG target=/root/build
ARG entrypoint=docker/docker-entrypoint.sh

RUN apk --no-cache add git && export GOPRIVATE=github.com/WhaleSu/wechatgpt && \
    export GOPROXY=https://goproxy.cn,direct

WORKDIR ${target}

COPY bootstrap/ ${target}/bootstrap
COPY config/ ${target}/config
COPY handler/ ${target}/handler
COPY local/ ${target}/local
COPY openai/ ${target}/openai
COPY utils/ ${target}/utils
COPY go.mod go.sum main.go ${target}/

RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GO111MODULE=on && go mod download && go build -o server main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /root/build/server .

CMD ./server

COPY ${entrypoint} /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT [ "docker-entrypoint.sh" ]
