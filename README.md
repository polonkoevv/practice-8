# Практическая работа №8
## Часть 1: gRPC
1) Установка необходимых инструментов
- Установили protoc
- Уставновили необходимые плагины для GRPC

2) Описание протокола в .proto файле
- Описали протокол с дополнительными методами (выборка по одному, выборка всех, удаление, создание, обновление)
```proto
syntax = "proto3";
package profile;

option go_package = "grpc-server/grpc";

service ProfileService {
  rpc GetProfile (ProfileRequest) returns (ProfileResponse);
  rpc CreateProfile (CreateProfileRequest) returns (ProfileResponse);
  rpc UpdateProfile (UpdateProfileRequest) returns (ProfileResponse);
  rpc DeleteProfile (ProfileRequest) returns (DeleteProfileResponse);
  rpc GetAllProfiles (EmptyRequest) returns (ProfileListResponse);
}

message ProfileRequest {
  int32 id = 1;
}

message CreateProfileRequest {
  string name = 1;
  string email = 2;
}

message UpdateProfileRequest {
  int32 id = 1;
  string name = 2;
  string email = 3;
}

message ProfileResponse {
  int32 id = 1;
  string name = 2;
  string email = 3;
}

message DeleteProfileResponse {
  bool success = 1;
  string message = 2;
}

message EmptyRequest {}

message ProfileListResponse {
  repeated ProfileResponse profiles = 1;
}

```

3) Генерация Go-кода
- Сгенерировали Go-код при помощи команды 
```bash 
protoc --go_out=. --go-grpc_out=. profile.proto
```
4) Реализация gRPC-сервера на Go
- Для хранения данных используется мапа со значениями типа Profile
```go 
type Profile struct {
	ID    int32
	Name  string
	Email string
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		profiles: make(map[int32]*pb.ProfileResponse),
		nextID:   1,
	}
}
```

### Команда для запуска gRPC сервера:
```bash 
go run grpc-server/main.go
```

5) Реализация gRPC-клиента

- Также сгенерировали код из того же proto файла.
- Реализовали в main.go использование всех методов.
### Команда для запуска теста:
```bash
go run grpc-client/main.go
```
### Примерный вывод теста:
```bash
Создан профиль: id:2  name:"Alice Johnson"  email:"alice@example.com"

Список всех профилей:
ID: 2, Name: Alice Johnson, Email: alice@example.com
ID: 1, Name: John Doe, Email: john@example.com

Получен профиль по ID: id:2  name:"Alice Johnson"  email:"alice@example.com"

Обновлен профиль: id:2  name:"Alice Smith"  email:"alice.smith@example.com"

Результат удаления: profile with id 2 successfully deleted

Список всех профилей после удаления:
ID: 1, Name: John Doe, Email: john@example.com
```

Сервер работает на порту 50051.
## Часть 2: GraphQL

1) Установка библиотеки gqlgen
- Установили требуемые библиотеки

2) Определение схемы
- Определили типы Profile и Company и требуемые для них CRUD методы
```graphql
type Profile {
  id: Int!
  name: String!
  email: String!
  companyID: Int
  company: Company
}

type Company {
  id: Int!
  name: String!
  address: String!
}
```

3) Реализация резолверов
- Реализовали резолверы для манипуляции сущностями.

4) Работа с аргументами, ошибками и вложенными типами
- Пример работы с аргументами - может кастомизировать вывод. Например, получим профили и выберем только id и имя профиля.
<image src="./images/customget.png" alt="Описание изображения">
[Получаем ошибку в выводе](/images/customget.png)

- Пример работы с аргументами - создание профиля с указанием id компании, которой не существует.
<image src="./images/grapherror.png" alt="Описание изображения">
[Получаем ошибку в выводе](/images/grapherror.png)

- Пример работы со вложенными типами - сначала создаем компанию, после этого создаем профиль, указав id нужной компании. При выборке профилей получаем:
<image src="./images/getallprofiles.png" alt="Описание изображения">
[Выборка всех профилей](/images/getallprofiles.png)

5) Запуск и регенерация
- Для регенерации кода по новой схеме нужно выполнить данную команду:
```bash 
cd graphql-api && go get github.com/99designs/gqlgen@v0.17.73 && go run github.com/99designs/gqlgen generate
```

- Для запуска сервера ужно выполнить данную команду:
```bash 
cd graphql-api && go run server.go
```

После запуска по адресу localhost:8080 станет доступен GUI для взаимодействия с API.
<image src="./images/gui.png" alt="Описание изображения">
[GUI](/images/gui.png)

## Примеры запросов к GraphQL
### Queries (Запросы)

### 1. Получить один профиль по ID
```graphql
query GetProfile {
  profile(id: 1) {
    id
    name
    email
    companyID
    company {
      id
      name
      address
    }
  }
}
```

### 2. Получить все профили
```graphql
query GetAllProfiles {
  profiles {
    id
    name
    email
    companyID
    company {
      id
      name
      address
    }
  }
}
```

### 3. Получить одну компанию по ID
```graphql
query GetCompany {
  company(id: 1) {
    id
    name
    address
  }
}
```

### 4. Получить все компании
```graphql
query GetAllCompanies {
  companies {
    id
    name
    address
  }
}
```

### Mutations (Мутации)

### 1. Создать профиль
```graphql
mutation CreateProfile {
  createProfile(input: {
    name: "John Doe"
    email: "john@example.com"
    companyID: 1
  }) {
    id
    name
    email
    companyID
    company {
      id
      name
      address
    }
  }
}
```

### 2. Обновить профиль
```graphql
mutation UpdateProfile {
  updateProfile(input: {
    id: 1
    name: "John Smith"
    email: "john.smith@example.com"
    companyID: 2
  }) {
    id
    name
    email
    companyID
    company {
      id
      name
      address
    }
  }
}
```

### 3. Удалить профиль
```graphql
mutation DeleteProfile {
  deleteProfile(id: 1)
}
```

### 4. Создать компанию
```graphql
mutation CreateCompany {
  createCompany(input: {
    name: "Tech Corp"
    address: "123 Tech Street"
  }) {
    id
    name
    address
  }
}
```

### 5. Обновить компанию
```graphql
mutation UpdateCompany {
  updateCompany(input: {
    id: 1
    name: "Tech Corp International"
    address: "456 Tech Avenue"
  }) {
    id
    name
    address
  }
}
```

### 6. Удалить компанию
```graphql
mutation DeleteCompany {
  deleteCompany(id: 1)
}
```

## Трудности
- При работе с graphql возниакли сложности с генерацией. При каждой новой генерации нужно устанавливать зависимость **go get github.com/99designs/gqlgen@v0.17.73**

## Сравнение REST vs gRPC vs GraphQL
REST:  
++ Простой и понятный  
++ Широкая поддержка инструментов  
++ Кэширование на уровне HTTP  
-- Избыточная передача данных (overfetching)  
-- Много запросов для получения связанных    данных (underfetching)  
Пример: GET /api/profiles/1/company  
gRPC:  
++ Высокая производительность (Protocol Buffers)  
++ Строгая типизация  
++ Двунаправленный стриминг  
++ Отлично подходит для микросервисов  
-- Сложнее в отладке  
-- Требует поддержки HTTP/2  
Пример: rpc GetProfile(ProfileRequest)  
returns (Profile)  
GraphQL:  
++ Получение только нужных данных  
++ Один endpoint для всех запросов  
++ Получение связанных данных в одном запросе  
-- Сложность на backend (N+1 проблема)  
-- Кэширование сложнее чем в REST  
Пример: { profile(id: 1) { name company { name } } }  
Когда что использовать:  
REST: для простых публичных API  
gRPC: для внутреннего взаимодействия микросервисов  
GraphQL: для сложных клиентских приложений с разными требованиями к данным