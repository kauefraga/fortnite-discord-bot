# Lau Discord Bot

> Convide-o para o seu servidor clicando [aqui](https://discord.com/api/oauth2/authorize?client_id=1201906543524851742&permissions=8&scope=bot).

## Comandos

- `!hello` - Responde com uma saudação
- `/pickgame` ou `!pickgame` - Responde com um jogo aleatório

## Como rodar

Pré-requisito: ter **o token do bot** em um arquivo `.env`.

### Desenvolvimento

Instale as dependências do projeto

```bash
go mod download
```

Rode o arquivo `cmd/main.go`

```bash
go run cmd/main.go
```

Certifique-se de que o arquivo `.env` esteja na raiz do projeto.

### Produção

Optei por criar uma imagem Docker e deixar o bot rodando em um container.

Primeiramente, baixe a imagem na sua máquina

```bash
docker pull kauefraga/lau-discord-bot
```

Agora, rode a imagem com o Docker

```bash
docker run -d --rm --name lau-discord-bot --env-file .env kauefraga/lau-discord-bot
```

A imagem do `Dockerfile` está disponível no Docker Hub: [kauefraga/lau-discord-bot](https://hub.docker.com/repository/docker/kauefraga/lau-discord-bot).

## Licença

Este projeto está sob licença do MIT - Veja a [LICENÇA](https://github.com/kauefraga/lau-discord-bot/blob/main/LICENSE) para mais informações.

---

Feito com ❤ por Kauê Fraga Rodrigues.
