FROM golang:1.23.4-alpine3.21 AS build

ARG APP_VERSION

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN go build -v -x -o pipebot-${APP_VERSION} cmd/pipebot/main.go

FROM alpine:3.21

ARG APP_VERSION
ENV APP_VERSION=${APP_VERSION}

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /app/pipebot-${APP_VERSION} .

CMD [ "sh", "-c", "./pipebot-${APP_VERSION}" ]
