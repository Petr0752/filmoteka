# 🎬 Filmoteka API

REST API для управления базой данных фильмов и актёров. Реализовано на Go с использованием Gin, Postgres и Swagger-документацией.

---

## 🚀 Возможности

- 🔐 Аутентификация с помощью JWT
- 👤 Ролевая модель (`admin`, `user`)
- 🎭 Управление актёрами (CRUD)
- 🎥 Управление фильмами (CRUD)
- 🔎 Поиск и сортировка фильмов
- 📑 Swagger-документация

---

## 📦 Технологии

- Go 1.22+
- Gin Web Framework
- PostgreSQL
- Swagger (`swaggo/swag`, `gin-swagger`)

---

## 🧪 Тестовые пользователи

База данных уже содержит два пользователя:

| Роль  | Логин  | Пароль     |
|-------|--------|------------|
| admin | admin  | admin123   |
| user  | user   | user123    |

Чтобы получить JWT токен:
1. Сделайте `POST /login` с телом:
   ```json
   {
     "username": "admin",
     "password": "admin123"
   }
