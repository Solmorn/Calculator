# Сервер для обработки арифметических выражений на Go

получает POST запросы и обрабатывает их




## запуск

go run ./cmd/main.go

## что есть в проекте

сервер,
тесты,
функция для вычисления

## возжможные ошибки

Код 200 - все хорошо.
Код 422 - в введенном выражении есть лишние символы, или оно неправильное.
Код 500 - ошибка выполнения подсчета (например деление на 0)



## примеры выражений


| выражение | вывод     | код                |
| :-------- | :------- | :------------------------- |
| 2+2 | 4 | 200 |
| 2+)2 | - | 422 |
| 2/(2-2) | - | 500 |