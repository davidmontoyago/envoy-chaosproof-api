# build
FROM golang:1.14.1-alpine3.11 AS build

RUN mkdir -p /app/bin

WORKDIR /app

COPY ./ /app/

RUN go mod vendor

RUN go build -i -v -o /app/bin/server /app

VOLUME /app

# run
FROM alpine:3.11

COPY --from=build /app/bin /app

CMD ["/app/server"]