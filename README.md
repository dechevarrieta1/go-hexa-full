# go-hexa-full

/internal
│   ├── /core
│   │   ├── user
│   │       ├── models/user.go       # Modelo User
│   │       ├── service/user_service.go # Lógica de negocio
│   ├── /ports
│   │   ├── user/user_repository.go  # Interfaz UserRepository
│   ├── /adapters
│   │   ├── mongo/mongo_user_repo.go # Adaptador MongoDB
│   ├── /infrastructure
│   │   ├── http/gin
│   │       ├── handlers/user_handler.go  # Handlers HTTP
│   │       ├── routes/router.go          # Rutas
│   │       ├── server.go                 # Servidor Gin
│   ├── /config
│       ├── config.go                    # Configuración


Ejemplo de como interactua esta arquiterctura

|-------------|        |---------------------|        |-------------------------|
| UserService | <----> | UserRepository Int  | <----> | MongoDB, Postgres, Mock |
|-------------|        |---------------------|        |-------------------------|


