FROM golang:1.23.4-alpine3.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/pipebot/main.go

FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY --from=build /app/main .

CMD [ "./main" ]
