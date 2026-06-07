# dmarket-data

CS2 skin data scraper and trading bot infrastructure.

## Structure

```
.
├── data/                 # Application source code
│   ├── bin/              # Config examples and data files
│   ├── cmd/              # CLI commands
│   ├── pkg/              # Shared packages (client, models, services, utils)
│   ├── migrations/       # Database migrations
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── docker-compose.yml    # Docker orchestration (NATS + app)
└── README.md
```

## Quick Start

```bash
# Start NATS and the application
docker compose up

```

## Data Sources

Skin and sticker definitions are sourced from the [CSGO-API](https://github.com/ByMykel/CSGO-API) repository.

## Setup

Copy example config files from [`data/bin/`](data/bin/):

```bash
cp data/bin/accs.example.json data/bin/accs.json
cp data/bin/skins_not_grouped.example.json data/bin/skins_not_grouped.json
cp data/bin/stickers.example.json data/bin/stickers.json
```

Copy and edit the config file:

```bash
cp data/bin/cfg.example.yaml data/bin/cfg.yaml
```

See [`data/README.md`](data/README.md) for full configuration details.