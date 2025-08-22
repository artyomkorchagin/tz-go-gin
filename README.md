# User API Service

REST API сервис для управления пользователями, разработанный на Go с использованием Gin framework.

## Технологии

- **Gin**
- **PostgreSQL**
- **Goose**
- **Zap**
- **Docker**
- **Viper**

## Makefile

| Команда        | Описание                           |
|----------------|------------------------------------|
| `make up`      | Запуск всех сервисов               |
| `make down`    | Остановка всех сервисов            |
| `make restart` | Перезапуск всех сервисов           |
| `make build`   | Сборка Docker образов              |
| `make db-up`   | Применение миграций БД             |
| `make db-down` | Откат миграций БД                  |
| `make clean`   | Полная очистка                     |

# Запуск
## 1. Клонирование репозитория
```
git clone <repository-url>
cd tz-go-gin
```

## 2. Запуск приложения
```
make up
```

# Использование
## Создание пользователя
```
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"login":"test_user","full_name":"Test User","gender":"male","age":25,"phone":"+1234567890","email":"test@example.com","avatar":"https://example.com/avatar.jpg","is_active":true}'
```
## Получение пользователя
```
curl -X GET http://localhost:3000/users/{user_id}
```
