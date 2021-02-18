# Blog REST API

## Build & Run (Locally)

- go 1.15
- docker
- golangci-lint (optional, used to run code checks)

Create .env file in root directory and add following values:

```bash
POSTGRES_USER=nikolas
POSTGRES_PASSWORD=admin
POSTGRES_PORT=5432
POSTGRES_HOST=localhost
POSTGRES_DB=blog
```

## Usage

`sudo docker-compose up --build` (Will create postgres db image and start http server)
