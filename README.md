# Auth Service (gRPC)

Простой gRPC-сервис для **аутентификации** пользователей.  
Хранит данные о пользователях, выполняет регистрацию и логин,  
выдаёт JWT-токен, который используют другие сервисы (например, [URL Shortener](https://github.com/Sheridanlk/UrlShortener).

Protobuf contract: https://github.com/Sheridanlk/protos

В дальнейшем планируется доработка до полноценного SSO-сервиса
