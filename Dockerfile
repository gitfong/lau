FROM golang:1.15.2-alpine3.12 AS build

LABEL maintainer="liujf1999@foxmail.com"

WORKDIR /usr/src

COPY ./ ./


RUN touch lau.log && chmod 666 lau.log

#RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN go build -o lau

RUN echo "#!/bin/sh" > lau.sh
RUN echo "./lau >>lau.log 2>&1" >> lau.sh 
RUN chmod +x lau.sh


FROM alpine:3.12
WORKDIR /usr/src

COPY --from=build /usr/src/lau* ./
COPY --from=build /usr/src/templates ./templates/
COPY --from=build /usr/src/resources ./resources/

EXPOSE 9090/tcp

#CMD ["param1"]
ENTRYPOINT ["./lau.sh"]
