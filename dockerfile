# build stage
FROM golang:1.12.6-alpine3.9 AS build-env
MAINTAINER wangkehenan <wangkehenan@gmail.com>
RUN apk add --no-cache git
ADD . /src
RUN cd /src && go build -o goapp cmd/debug-unhealthy/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp