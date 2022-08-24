# deploy-builder
FROM golang:1.19.0-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# deploy
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD [ "./app" ]

# dev
FROM golang:1.19.0-bullseye as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD [ "air" ]

