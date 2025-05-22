# Цитатник — мини-сервис на Go


## 🚀 Как запустить

1. Клонируйте репозиторий:

```bash
git clone https://github.com/your-username/quotes-api.git
cd quotes-api
```

3. Запустите сервер:

```bash
go run main.go
```

4. Сервер будет работать на порту `http://localhost:8080`

---

## 📌 Поддерживаемые эндпоинты

| Метод | Эндпоинт                     | Описание                                  |
|-------|------------------------------|-------------------------------------------|
| POST  | `/quotes`                   | Добавить новую цитату                     |
| GET   | `/quotes`                   | Получить все цитаты                       |
| GET   | `/quotes/random`           | Получить случайную цитату                 |
| GET   | `/quotes?author=Автор`     | Фильтровать цитаты по автору             |
| DELETE| `/quotes/{id}`             | Удалить цитату по ID                      |

---

## 🔧 Примеры curl-команд

### ➕ Добавить цитату

```bash
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d "{\"author\":\"Confucius\", \"text\":\"Life is simple, but we insist on making it complicated.\"}"
```

### 📄 Получить все цитаты

```bash
curl http://localhost:8080/quotes
```

### 🎯 Получить случайную цитату

```bash
curl http://localhost:8080/quotes/random
```

### 🔍 Фильтрация по автору

```bash
curl "http://localhost:8080/quotes?author=Confucius"
```

### ❌ Удалить цитату по ID

```bash
curl -X DELETE http://localhost:8080/quotes/1
```

---

## 🧪 Тесты

> (опционально) Можно добавить unit-тесты с использованием стандартной библиотеки `testing`.

---

## 🛠 Используемые технологии

- Язык программирования: Go
- Роутер: `gorilla/mux`
- Хранение данных: in-memory (slice)

---