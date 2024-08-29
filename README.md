# todo_list_golang

Цель: Разработать REST API для системы управления задачами, которая позволяет пользователям создавать, просматривать, обновлять и удалять задачи.
Проект релизован с помощью Gin, Gorm, Postgres, Goose, Zerolog, Redis, Docker, Docker-compose

## Чтобы запустить проект у себя локально, вам необохимо
Склонировать репозиторий

`git clone git@github.com:fazletdinov/todo_list_golang.git`

Скопируйте содержимое .env.example в .env

И запустите команду

```
make start
```
Если не установлена утилита make, то необходимо запустить следующей командой

```
docker compose -f docker-compose.yaml up --build
```
Вышеуказанные команды запустит приложение
Далее можете посмотреть Api спецификацию (свагер) по адресу:
`http://localhost:8000/docs/index.html`

### Для создания задачи

`POST api/v1/tasks` вернет созданный объект Task.

### Пример тела запроса
```json
{
  "description": "Description",
  "due_date": "02 Jan 2025 05:55PM",
  "title": "Title"
}
```

### Ответ

Успешный ответ приходит с кодом `201 Created` и содержит тело:

```json
{
  "id": 2,
  "title": "Title",
  "description": "Description",
  "due_date": "2 Jan 2025 5:55PM",
  "created_at": "2024-08-29T21:38:58.36030694+08:00",
  "update_at": "2024-08-29T21:38:58.36030694+08:00"
}
```

### Для получения списка задач
`GET api/v1/tasks` возвращает список задач.

Параметры:
* `limit` – количество
* `offset` – смещение

### Пример запроса
`GET api/v1/tasks?limit=10&offset=0`

### Ответ
Успешный ответ приходит с кодом `200 OK` и содержит тело:
```json
{
[
  {
    "id": 2,
    "title": "Title",
    "description": "Description",
    "due_date": "2 Jan 2025 5:55PM",
    "created_at": "2024-08-29T21:38:58.360306+08:00",
    "update_at": "2024-08-29T21:38:58.360306+08:00"
  },
  {
    "id": 3,
    "title": "Title 2",
    "description": "Description 2",
    "due_date": "23 Sep 2024 5:55PM",
    "created_at": "2024-08-29T21:44:02.556942+08:00",
    "update_at": "2024-08-29T21:44:02.556942+08:00"
  }
]
}
```

### Для получения задачи по id
`GET api/v1/tasks/{id}/` возвращает задачу.

### Пример запроса
`GET api/v1/tasks/2/`

### Ответ
Успешный ответ приходит с кодом `200 OK` и содержит тело:
```json
{
    "id": 2,
    "title": "Title",
    "description": "Description",
    "due_date": "2 Jan 2025 5:55PM",
    "created_at": "2024-08-29T21:38:58.360306+08:00",
    "update_at": "2024-08-29T21:38:58.360306+08:00"
}
```

### Для обновления задачи по id
`PUT api/v1/tasks/{id}/` возвращает обновленную задачу.

### Пример тела запроса
```json
{
  "description": "Description two",
  "due_date": "7 Jan 2025 5:06PM",
  "title": "Title two"
}
```

### Ответ
Успешный ответ приходит с кодом `200 OK` и содержит тело:
```json
{
  "id": 2,
  "title": "Title two",
  "description": "Description two",
  "due_date": "7 Jan 2025 5:06PM",
  "created_at": "2024-08-29T21:56:36.997353+08:00",
  "update_at": "2024-08-29T21:57:49.504185+08:00"
}
```
### Для Удаления задачи по id
`DELETE api/v1/tasks/{id}/` .

### Ответ
Успешный ответ приходит с кодом `204 NO CONTENT` и без тела

### Автор
[Idel Fazletdinov - fazletdinov](https://github.com/fazletdinov)