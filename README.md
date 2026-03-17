# apex20-ws

Serviço de WebSocket do **Apex20** — Virtual Tabletop para RPG.

Construído em Go com Gorilla WebSocket. Inscrito no Redis Pub/Sub para receber eventos da API e fazer broadcast para os clientes conectados em tempo real.

## Pré-requisitos

- Go v1.26+
- Redis v7

## Instalação

```bash
# Ativar git hooks locais
make setup

# Copiar e preencher variáveis de ambiente
cp .env.example .env
```

## Comandos

```bash
make run     # inicia o servidor WebSocket
make build   # compila o binário em bin/ws
make test    # executa os testes
make lint    # golangci-lint
```

## Estrutura

```
cmd/ws/                               Entrypoint
internal/
  application/port/                   Interfaces (ports)
  infrastructure/adapter/inbound/     Handler WebSocket
  infrastructure/adapter/outbound/    Subscriber Redis
```

## Variáveis de ambiente

Copie `.env.example` para `.env` e preencha os valores.

## Documentação

Consulte o submodule `docs/` ou o repositório [apex20-docs](https://github.com/lins-dev/apex20-docs).
