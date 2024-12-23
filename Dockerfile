FROM golang:1.23.4-alpine3.21 AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build cmd/pipebot/main.go

FROM golang:1.23.3-alpine3.20

WORKDIR /app

COPY --from=build /app/main .

CMD [ "./main" ]
