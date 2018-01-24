## Dependencies

We are using [dep](https://github.com/golang/dep) for dependency management. 

## Installing

```sh
go get -u github.com/golang/dep/cmd/dep
```

## Restoring

```sh
dep ensure
```

You are ready to Go :)

[See more about dep here](https://github.com/golang/dep)

## Auth0

All request will the authenticated against **Auht0**.
If you wanna disable it, just remove `v1.Use(middlewares.Authentication)`.

You can configure your [**Auth0**](https://auth0.com/) settings in `.env` file.

```
AUTH0_JWKS=https://{YOUR_AUTH0_TENANT_DOMAIN}/.well-known/jwks.json
AUTH0_AUDIENCE=http://{{YOUR_AUTH0_API_DOMAIN}
AUTH0_ISSUER=https://{YOUR_AUTH0_TENANT_DOMAIN}.auth0.com/
```

## Swagger

### What is Swagger?

Swagger is a set of open-source tools built around the OpenAPI Specification that can help you design, build, document and consume REST APIs.

For use Swagger in this project, you'll need *gin-swagger*.

### Installing

Install `swag` using `go get`

```sh
go get github.com/swaggo/swag/cmd/swag
```

### Generating docs

You can generate or regenerate your docs using the command below:

```sh
swag init
```

[See more about swag here](https://github.com/swaggo/gin-swagger)

## Environment

`.env` file should have all variables you'll need on your environment.
`godotenv` will read your `.env` file from the project folder or `projetc/secrets` folder and create the envs.

[See more about godotenv here](https://github.com/joho/godotenv)

## Docker

You can run this code easily on a Docker container.
Just run the following commands:

```sh
docker build -t your-image .
docker run -p 9000:9000 your-image
```
