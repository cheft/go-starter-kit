FROM alpine:3.9

# 设置时区为上海
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

ADD app /go/bin/
WORKDIR /go/bin

CMD ["/go/bin/app"]

ARG APPLICATION_PORT
EXPOSE ${APPLICATION_PORT}
