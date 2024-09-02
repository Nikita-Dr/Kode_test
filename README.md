# API Documentation

## Endpoints

### SignUp
Регистрация нового пользователя
- **Endpoint:**  /auth/signup
- **Method:** POST

Необходимо передать логин и пароль. Поле email не валидируется.
```json
{
  "email": "testemail"
  "password": "password"
}
```
В случае успеха вернет статус 200


### Login
Авторизация пользователя
- **Endpoint:**  /auth/login
- **Method:** POST
  
Необходимо передать логин и пароль. 
```json
{
  "email": "testemail"
  "password": "password"
}
```
В случае успеха вернет jwt token в json
```json
{
  "data": "token"
}
```


---

### Get notes
Получает заметки пользователя
- **Endpoint:**  /notes
- **Method:** GET
  
Необходимо передать jwt token в заголовке запроса в поле "Authorization"
В случае успеха вернет список заметок
```json
{
  "id": "id"
  "note": "note text"
}
```

### Create note
Создает заметку
- **Endpoint:**  /notes
- **Method:** POST
  
Необходимо передать jwt token в заголовке запроса в поле "Authorization" и заметку в виде json
```json
{
  "id": "id"
  "note": "note text"
}
```

## Яндекс Спеллер
При создании заметки на слое usecase internal/domain/note/usecase/note.go вызывается метод ValidateText(note.Note) куда передается текст заметки. 
Далее в файел internal/domain/note/usecase/text_checker.go текст валидируется и передается в pkg/yadex/yandex.go где происходит запрсос к Яндекс Спеллеру.
Для проверки текста берется первые предложенные варианты слов от Яндекса и затем передаются в бд.
