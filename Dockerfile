FROM golang:1.21-alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags=-w -o /bin/lau-discord-bot ./cmd/main.go

FROM alpine AS production

RUN apk --no-cache add ca-certificates

COPY --from=build /bin/lau-discord-bot /bin/lau-discord-bot

ENTRYPOINT [ "/bin/lau-discord-bot" ]
