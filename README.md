# HTTP-API - сервис

#Для запуска сервера - go run cmd/main.go

В корневой папке имеется Postman Collections для проведения тестов.

#Docker - postgresSQL

1) sudo docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
2) migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up (первый запуск)
3) sudo docker exec -it <conteiner> /bin/bash
4) psql -U postgres

#API

go run cmd/main.go

#Request

1. Добавить нового пользователя - curl --location --request POST 'localhost:8080/api/createUser' \
   --header 'Content-Type: application/json' \
   --data-raw '{
   "name": "Eremei",
   "last_name": "Kravcov",
   "phone": 33333
   }'
2. Получить список всех пользователей - curl --location --request GET 'localhost:8080/api/users'
3. Удалить конкретного пользователя по id - curl --location --request GET 'localhost:8080/api/user/2'

