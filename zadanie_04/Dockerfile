FROM golang:1.24.2-alpine AS build

WORKDIR /build

RUN apk add build-base

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o app cmd/main.go

FROM golang:1.24.2-alpine

WORKDIR /app

COPY --from=build build/app .
COPY --from=build build/db ./db/

ENTRYPOINT ["./app"]
