# BirthdayGreetingService

Запрос для создание списка пользователей 
```bash
curl -b cookies.txt -X POST http://localhost:8080/auth/user
-H "Content-Type: application/json"
-d '{
      "name":"Ivan",
      "sur_name":"Bober",
      "birth_day":"2002-06-11"
    }'
```
