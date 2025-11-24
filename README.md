# Auth Service (gRPC)

Простой gRPC-сервис для **аутентификации** пользователей.  
Хранит данные о пользователях, выполняет регистрацию и логин,  
выдаёт JWT-токен, который используют другие сервисы (например, [URL Shortener](https://github.com/Sheridanlk/UrlShortener)).

Protobuf contract: https://github.com/Sheridanlk/protos

---

# 1. Технологический стек

| Категория       | Технологии                       |
| --------------- | -------------------------------- |
| Транспорт       | gRPC + Protocol Buffers          |
| Auth            | JWT (HS256), bcrypt              |
| БД              | PostgreSQL                       |
| Миграции        | **golang-migrate/migrate**       |
| Конфигурация    | cleanenv + YAML/ENV              |
| Логирование     | slog                             |
| Тесты           | testify, mockery                 |
| Контейнеризация | Docker + Docker Compose          |

---

# 2. Основной функционал

- Регистрация пользователя (проверка уникальности, хеширование пароля — bcrypt)

- Логин и выдача JWT (issuer/audience/ttl/secret из конфига)

- Хранение пользователей в PostgreSQL

- Миграции схемы через golang-migrate

- Логи на slog, конфигурация через cleanenv + YAML

---

# 3. Структура

```
gRPCAuth/
├── cmd/
│   └── sso/                 # точка входа
├── internal/
│   ├── config/               # cleanenv, структура конфигурации
│   ├── app/                  # application layer
│   ├── logger/               # настройка slog
│   ├── server/               # gRPC-сервер, регистрация сервисов/интерсепторов
│   ├── service/              # бизнес-логика: Register/Login/Validate
│   ├── storage/
│   │   └── postgresql/       # репозиторий пользователей (pgx/sqlx)
│   ├── auth/                 # JWT: генерация, проверка, claims
│   └── lib/                  # утилиты/хелперы (ошибки, валидация и т.п.)
├── migrations/               # SQL миграции для golang-migrate
├── config/
│   └── config.yaml           # пример конфигурации
├── Makefile                  # сборка, запуск, миграции, docker
├── Dockerfile
└── docker-compose.yml

```

---

# 4. Запуск
## 4.1. Конфигурация
Создать файл конфигурации config/config.yaml:
```yaml
env: "local" #dev, prod
token_ttl: 1h
postgres:
  host: auth-db 
  port: 5432
  user: postgres
  password: postgres
  name: auth
grpc:
  port: 34443
  timeout: 20s
```
(Для локального запуска без Docker: поставить host: localhost и реальный порт)

## 4.2. Старт

A) Локально(Go):
 
```
make run-local
```

B) В Docker контейнере(рекомендуется)

```
make run-docker
```

## 4.3 Миграции

Создать новую:

```
make create-migration NAME=<имя_миграции>
```

Поднять в контейнере:

```
make migrate-up-docker    # применить все
make migrate-down-docker  # откатить одну
```

Локально(поменять путь для подключения к бд на свой)

```
make migrate-up-local    # применить все
make migrate-down-local  # откатить одну
```

# 5. TODO
В дальнейшем планируется доработка до полноценного SSO-сервиса, с использованием OAuth 2.0
