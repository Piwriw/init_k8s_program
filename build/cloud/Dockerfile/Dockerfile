FROM golang:1.17-alpine3.14 AS builder

COPY . /home/code/uavedge

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /home/code/uavedge

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk add --no-cache gcc musl-dev


RUN CGO_ENABLE=0 go build -o /usr/local/bin/agent-cloud /home/code/uavedge/cloud/cmd

FROM alpine:3.14

COPY --from=builder /usr/local/bin/agent-cloud /usr/local/bin/agent-cloud

ENTRYPOINT ["agent-cloud"]