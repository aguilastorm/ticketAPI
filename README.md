# TicketAPI

Este proyecto es una API REST en Go (Golang) para gestionar un sistema de tickets. El objetivo principal es proporcionar un servicio que permita a los usuarios crear, leer, actualizar y eliminar tickets.

## Estructura del Proyecto

El proyecto se estructura de la siguiente manera:

    .
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── handlers
    │   ├── createTicket.go
    │   ├── deleteTicket.go
    │   ├── getTicket.go
    │   ├── getTickets.go
    │   └── updateTicket.go
    ├── models
    │   └── ticket.go
    ├── store
    │   └── store.go
    └── tests
        ├── createTicket_test.go
        ├── deleteTicket_test.go
        ├── getTicket_test.go
        ├── getTickets_test.go
        └── updateTicket_test.go


Las carpetas son las siguientes:

- `handlers/`: Contiene los controladores para cada operación CRUD (Crear, Leer, Actualizar, Eliminar). Cada controlador tiene su propio archivo.
- `models/`: Contiene la definición del modelo de datos para un Ticket.
- `store/`: Contiene la lógica para almacenar y recuperar Tickets. En esta versión simple, los Tickets se almacenan en memoria.
- `tests/`: Contiene pruebas unitarias para cada controlador.

## Instalación y Uso

Para instalar y ejecutar el proyecto, sigue estos pasos:

1. Clona el repositorio en tu máquina local.
2. Navega hasta el directorio del proyecto en tu terminal.
3. Ejecuta `go mod tidy` para descargar las dependencias necesarias.
4. Ejecuta `go run main.go` para iniciar el servidor.

La API estará disponible en `http://localhost:8080`.

## Pruebas

Para ejecutar las pruebas unitarias, usa el comando `go test ./...` en el directorio raíz del proyecto.

## Endpoints

Los endpoints disponibles son los siguientes:

- `GET /tickets`: Devuelve todos los tickets.
- `GET /tickets/{id}`: Devuelve un ticket específico.
- `POST /tickets`: Crea un nuevo ticket.
- `DELETE /tickets/{id}`: Elimina un ticket específico.
- `PUT /tickets/{id}`: Actualiza un ticket específico.
