FROM golang:alpine AS base_alpine

WORKDIR /go/src

ADD . /go/src

WORKDIR /go/src/ChitChat

RUN cd /go/src/ChitChat && go build -i -v -o ChitChat && mkdir /final \
    && cp -r /go/src/ChitChat/ChitChat /final \
    && cp -r /go/src/ChitChat/public /final \
    && cp -r /go/src/ChitChat/templates /final \
    && cp -r /go/src/ChitChat/config.json /final \
    && cp -r /go/src/ChitChat/chitchat.log /final \
    && cp -r /go/src/ChitChat/config.json /final

FROM alpine

RUN apk update && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /myApp

COPY --from=base_alpine /final /myApp

EXPOSE 8080

ENTRYPOINT ./ChitChat