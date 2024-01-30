FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/fortnite-discord-bot ./cmd/main.go

ENTRYPOINT [ "/bin/fortnite-discord-bot" ]
