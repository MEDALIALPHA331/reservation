# TODOs

- sqlx, sqlite & turso db
- auth
- testing & coverage
- other CRUD
- config
- docker
- gh actions, ci/cd
- pulumi, iac -lambda
- frontend: templates, go temp, htmx + Vite

## enviroment variables

- First install godotenv package

  ```bash
  go get github.com/joho/godotenv
  ```

- auto import env variables from .env in main.go using:

  ```go
  import _ "github.com/joho/godotenv/autoload"
  ```

## Install Echo http framework

- ```bash
    go get github.com/labstack/echo/v4
    go get github.com/labstack/echo/v4/middleware
  ```

## Mongodb docker image

- https://hub.docker.com/_/mongo
- open docker and run this command to start mongodb:

  ```bash
  docker run -p 27017:27017 mongo:latest
  ```

## Mongodb compass tool

- [Download Compass from here](https://www.mongodb.com/try/download/compass)

## mongodb go driver docs

- [visit this link for the mongodb go driver docs](https://www.mongodb.com/docs/drivers/go/)

## mongodb go driver package:

- [Visit this link to download the driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)

```bash
go get go.mongodb.org/mongo-driver/mongo
```

- Install [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt#GenerateFromPassword):

  ```bash
    go get golang.org/x/crypto/bcrypt
  ```

- To query api using curl e.g

```bash
    curl -v -X POST -H "Content-Type: application/json" -d '{"FirstName": "salah", "LastName": "khaled", "Email": "salah@gmail.com", "Password": "dali12346789"}' http://127.0.0.1:8080/apiv1/user

```
