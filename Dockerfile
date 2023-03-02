#build stage
FROM golang:alpine AS builder

LABEL maintainer="uerax"

WORKDIR /app
COPY . /app

ENV GOPROXY https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o chatgpt1 main.go

FROM alpine

WORKDIR /app

# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /app/chatgpt1 .
COPY --from=builder /app/etc ./etc

# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata

EXPOSE 8080

VOLUME ["/app/etc","/app/log"]

ENTRYPOINT [ "./chatgpt1" ]
