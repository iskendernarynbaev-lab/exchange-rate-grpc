# exchange-rate-grpc

Простой и production-ready gRPC сервис на Go, который:
- забирает стакан USDT с Grinex (`/api/v1/spot/depth?symbol=usdta7a5`)
- считает `ask` и `bid` (по `topN` или `avgNM`)
- сохраняет каждый результат `GetRates` в PostgreSQL
- отдает healthcheck, метрики Prometheus и трассировку OpenTelemetry

## Что внутри

- gRPC метод: `rates.v1.RatesService/GetRates`
- gRPC healthcheck: `grpc.health.v1.Health/Check`
- Логирование: `zap`
- Метрики: `/metrics`
- Graceful shutdown
- Миграции: `goose` (вручную, не автоматически)

## Структура

```text
cmd/app                      # точка входа
internal/app                 # запуск приложения (orchestration)
internal/client/grinex       # HTTP клиент Grinex
internal/config              # .env + yaml + флаги
internal/grpcserver          # gRPC handlers
internal/repository/postgres # работа с БД
internal/service             # бизнес-логика
internal/storage/migrations  # goose миграции
pkg/api/rates/v1             # protobuf
configs/config.yaml          # конфиг приложения
.env                         # локальные переменные
```

## Требования

- Go `1.25+`
- Docker + Docker Compose

## Быстрый старт

```bash
make docker
```

`make docker` (или `make docker-build`) выполняет шаги последовательно:
1. поднимает PostgreSQL
2. запускает миграции (`migrator`)
3. поднимает `app`

Проверка метода (`topn`, взять 1-й уровень):

```bash
grpcurl -plaintext -d '{"method":"topn","n":1}' localhost:9090 rates.v1.RatesService/GetRates
```

Проверка метода (`avgnm`, среднее по диапазону 1..3):

```bash
grpcurl -plaintext -d '{"method":"avgnm","n":1,"m":3}' localhost:9090 rates.v1.RatesService/GetRates
```

Проверка healthcheck:

```bash
grpcurl -plaintext -d '{"service":""}' localhost:9090 grpc.health.v1.Health/Check
```

## Примеры gRPC запросов

`topn` (взять N-ю позицию стакана):

```bash
grpcurl -plaintext -d '{"method":"topn","n":1}' localhost:9090 rates.v1.RatesService/GetRates
```

`avgnm` (среднее по диапазону [N;M]):

```bash
grpcurl -plaintext -d '{"method":"avgnm","n":1,"m":3}' localhost:9090 rates.v1.RatesService/GetRates
```

Пример невалидного запроса (`m < n` для `avgnm`, вернется `InvalidArgument`):

```bash
grpcurl -plaintext -d '{"method":"avgnm","n":5,"m":3}' localhost:9090 rates.v1.RatesService/GetRates
```

## Команды Make

```bash
make build
make test
make run
make docker-build
make docker
make lint

make migrate-status
make migrate-up
make migrate-down
make migrate-create NAME=add_new_table
```

`make docker-build` (или коротко `make docker`) поднимает весь стек в правильном порядке (PostgreSQL -> migrator -> app).

## Конфиг

Порядок приоритета:
1. флаги запуска
2. переменные окружения (`.env`)
3. `configs/config.yaml`

По умолчанию используется `configs/config.yaml`.
Можно переопределить:
- `CONFIG_PATH=/path/to/config.yaml`
- `-config /path/to/config.yaml`

Пример плейсхолдера в YAML:

```yaml
server:
  grpc_addr: ${GRPC_ADDR::9090}
  shutdown_timeout: ${SHUTDOWN_TIMEOUT:10s}
```

## Основные переменные

- `DATABASE_URL` — строка подключения к PostgreSQL
- `GRPC_ADDR` — адрес gRPC (по умолчанию `:9090`)
- `METRICS_ADDR` — адрес HTTP метрик (по умолчанию `:2112`)
- `GRINEX_URL`, `GRINEX_SYMBOL` — источник котировок
- `LOG_LEVEL` — `debug|info|warn|error`

## Миграции

Миграции лежат в:

- `internal/storage/migrations`

Можно запускать вручную через `make migrate-*`.
При `make docker-build` миграции запускаются автоматически шагом `docker-compose run --rm migrator`.

## Метрики и health

- `http://localhost:2112/metrics`
- `http://localhost:2112/healthz`
