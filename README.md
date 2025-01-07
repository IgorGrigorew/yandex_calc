Онлайн калькулятор

Калькулятор выполняет простые матиматические операции, которые берет в json формате из POST запроса пользователя.
Обрабатыват ошибки

- Ошибка 422: Expression is not valid.
- Ошибка 500: Internal server error.


  
Пример запроса пользователя
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

Запуск 
go run ./cmd/main.go
