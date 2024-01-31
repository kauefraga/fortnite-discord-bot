FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/lau-discord-bot ./cmd/main.go

ENTRYPOINT [ "/bin/lau-discord-bot" ]
