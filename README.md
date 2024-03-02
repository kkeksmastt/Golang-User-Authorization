# Golang-User-Authorization

### Test task BackDev

Тестовое задание на позицию Junior Backend Developer

**Используемые технологии:**

- Go
- JWT
- MongoDB

**Задание:**

Написать часть сервиса аутентификации.

Два REST маршрута:

- Первый маршрут выдает пару Access, Refresh токенов для пользователя сидентификатором (GUID) указанным в параметре запроса
- Второй маршрут выполняет Refresh операцию на пару Access, Refreshтокенов

**Требования:**

Access токен тип JWT, алгоритм SHA512, хранить в базе строго запрещено.

Refresh токен тип произвольный, формат передачи base64, хранится в базеисключительно в виде bcrypt хеша, должен быть защищен от изменения настороне клиента и попыток повторного использования.

Access, Refresh токены обоюдно связаны, Refresh операцию для Access токена можно выполнить только тем Refresh токеном который был выдан вместе с ним.

## Описание решения

### Указание переменных

Для указания порта сервера(по умолчанию 8000) и MongoDB(по умолчанию mongodb://localhost:27017) использовался файл .env

Указать время жизни токена возможно в переменной LiveTimeOfToken, а метод шифрования в SigningMethod (UserAuth/jwt_token/token.go)

Пакеты указаны по пути UserAuth/{name}, пожалуйста, учитывайте это при запуске кода

### Тест работы

Для тесторования использовался Postman (www.postman.com/downloads/)

1. Тестовый Guid, отправлялся по POST localhost:8000/api/get-token:
```
{
    "guid":"7F9619FF-8B86-D011-B42D-00CF4FC964FF"
}
```
2. Ответ от сервера (access и refresh токены) отправлялись по PUT localhost:8000/api/refresh-token:
```
{
"access":"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.
          eyJndWlkIjoiN0Y5NjE5RkYtOEI4Ni1EMDExLUI0MkQtMDBDRjRGQzk2NEZGIiwiZXhwIjoxNzEyOTIxNTQyfQ.
          CE6zK7_VTa_-5_sushY5vHv1YE5FVbTGDMi0pbaEji-Goc5Q8kgT3d23clO9ZIknusEl585vMg9eLLjkieD-lg",
"refresh": "V1dXV1dXV1dXVw=="
}
```
(Не стоит копировать данные токены. Лучше взять свежие из ответа сервера)

### Color

Пакет добавлен в целях покраски логов сервера, InfoLog будет окрашен в зеленый, а Error в красный (на windows стандартный)
