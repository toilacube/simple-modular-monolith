## Go tutorial 

A simple member/movie CRUD project but in a little bit complex way to demonstrate how to build a modular monolith application in Go:
 - Member can register, login, update profile and delete account
 - Movie can be created, read, updated and deleted by member

 ## Architecture
 - Follow Modular Monolith architecture:
    - Each module has its own layer (handler, service, repository).
    - Each module is independent and can be developed/tested separately.

- Including 3 modules:
    - Member module
    - Movie module
    - Auth module

- A gateway will sit in front of all modules to route the request to the correct module using Caddy server.


## API
| Method | Endpoint | Description | Response Type | Response |
| :--- | :--- | :--- | :--- | :--- |
| `GET` | `/ping` | Check status server | `plaintext` | `pong` |
| `POST` | `/api/v1/member/register` | Register member | `plaintext` | `No content` |
| `POST` | `/api/v1/member/login` | Login user, return JWT | `application/json` | ```json { "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30" } ``` |
| `POST` | `/api/v1/movies` | Create new movie | `application/json` | ```json { "id": "632099a6-a697-4cbf-8e79-298ae9d7997d", "name": "Star War", "star": 5, "actor": "Mark Hamill", "created_at": 1758092201, "updated_at": 1758092201 } ``` |
| `GET` | `/api/v1/movies` | Get all of movies created by creator. Only creator can call this endpoint. | `application/json` | ```json { "creator": "3053dd78-b24e-4ff5-a3c3-53098982aed9", "name": "Jhon Doe", "movies": [{ "id": "632099a6-a697-4cbf-8e79-298ae9d7997d", "name": "Star War", "star": 5, "actor": "Mark Hamill", "created_at": 1758092201, "updated_at": 1758092201 }] } ``` |