# dmarket-bot

CS2 skin data scraper that aggregates listing data from dmarket and Buff.

## Architecture

```
cmd/          # CLI entry points (run, config, subcommands)
pkg/          # Shared libraries
├── client/   # API clients (dmarket, buff)
├── models/   # Domain models (Sale, Item, Skin, etc.)
├── services/ # Business logic (scraper)
└── utils/    # Helpers
bin/          # Config files and data (cfg.yaml, accs.json)
migrations/   # DB schema (unused - sqlc configured but not used)
```

## Usage

```bash
# Run all enabled scrapers (uses config defaults)
go run main.go run --cfg ./bin/cfg.yaml

# Override delay for a specific module
go run main.go run --cfg ./bin/cfg.yaml --dmarket.new_listing.delay=200ms

# Check available flags
go run main.go run --help
```

## Configuration

### bin/cfg.yaml

Copy from example:

```bash
cp bin/cfg.example.yaml bin/cfg.yaml
```

Controls which scrapers run and their settings:

```yaml
Dmarket:
  NewListing:
    enabled: true
    acc_path: "./bin/accs.json"
    delay: 100ms
  OrderBook:
    enabled: true
    acc_path: "./accs.json"
    skins_path: "./bin/skins_not_grouped.json"
    stickers_path: "./bin/stickers.json"
    delay: 100ms
Buff:
  Listing:
    enabled: true
    acc_path: "./accs.json"
    mapping_path: "./bin/cs2_marketplaceids.json"
    delay: 100ms
```

### CLI Flags (override config)

| Flag | Default | Description |
|------|---------|-------------|
| `--cfg` | (required) | Path to config YAML |
| `--dmarket.new_listing.delay` | 100ms | New listing poll interval |
| `--dmarket.orderbook.delay` | 100ms | Order book poll interval |
| `--buff.listing.delay` | 100ms | Buff listing poll interval |

Flags take precedence over config values when explicitly provided.

### bin/accs.json

Account credentials for dmarket API. Copy from example:

```bash
cp bin/accs.example.json bin/accs.json
```

Each entry supports:
- `secret_key` - DMarket API secret key
- `private_key` - RSA private key for authentication
- `proxy_url` - HTTP proxy URL (optional)

### bin/skins_not_grouped.json

CS2 skin definitions. Copy from example:

```bash
cp bin/skins_not_grouped.json.example bin/skins_not_grouped.json
```

Each entry represents a skin variant with wear conditions, float ranges, and metadata.

### bin/stickers.json

CS2 sticker definitions. Copy from example:

```bash
cp bin/stickers.json.example bin/stickers.json
```

Each entry represents a sticker with rarity and metadata.

### bin/cs2_marketplaceids.json

Marketplace ID mappings (required for BUFF order book scraping).

### Lookup Configuration

The `Lookup` section in `cfg.yaml` defines paths to all CS2 data files used to build the item lookup table:

```yaml
Lookup:
  skins_path: "./bin/skins_not_grouped.json"
  stickers_path: "./bin/stickers.json"
  agents_path: "./bin/cs2/agents.json"
  graffiti_path: "./bin/cs2/graffiti.json"
  crates_path: "./bin/cs2/crates.json"
  keychains_path: "./bin/cs2/keychains.json"
  music_kits_path: "./bin/cs2/music_kits.json"
  patches_path: "./bin/cs2/patches.json"
  sticker_slabs_path: "./bin/cs2/sticker_slabs.json"
  collectibles_path: "./bin/cs2/collectibles.json"
```

Items are keyed by `market_hash_name` (falling back to `name` when unavailable).

### bin/cs2/

CS2 item definitions sourced from [CSGO-API](https://github.com/ByMykel/CSGO-API):

| File | Description |
|------|-------------|
| `agents.json` | Agent characters with rarity and team info |
| `crates.json` | Weapon cases, souvenir packages, and capsules |
| `graffiti.json` | Spray patterns and designs |
| `keychains.json` | Weapon charms |
| `music_kits.json` | Music kits with MVP anthems |
| `patches.json` | Agent patches |
| `sticker_slabs.json` | Sealed sticker packs |
| `collectibles.json` | Commemorative coins, medals, and trophies |

### Data Sources

Skin, sticker, and all CS2 item definitions can be obtained from the [CSGO-API](https://github.com/ByMykel/CSGO-API) repository, which provides comprehensive CS2 item definitions.

## Running with Docker

```bash
# From project root
docker compose up
```

The app connects to a NATS server running in the same compose network.

## Dependencies

- [NATS](https://nats.io/) - Message broker (in-container or external)
- [cobra](https://github.com/spf13/cobra) - CLI framework
- [zap](https://github.com/uber-go/zap) - Structured logging
