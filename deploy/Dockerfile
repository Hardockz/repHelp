FROM alpine:3.8

RUN apk --update --no-cache add ca-certificates tzdata curl shadow \
    && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
    && echo "America/Sao_Paulo" > /etc/timezone \
    && mkdir -p /home/default \
    && groupadd -r default \
    && useradd --no-log-init -r -g default default 

COPY build/mattermost-bot-test /usr/bin

USER default

CMD mattermost-bot-test