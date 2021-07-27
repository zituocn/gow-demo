FROM alpine:latest

# app name
ENV AppName=grpc-client




# copy file to image
COPY ${AppName} /app/


# add timezone
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories \
        && apk --no-cache add ca-certificates tzdata\
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone \
        && apk del tzdata

WORKDIR /app


# port
EXPOSE 9090

# label
LABEL maintainer = "301109640@qq.com"

CMD ["sh","-c","./$AppName"]