## Dependencies

`dep` is a dependency management Go. 

```sh
go get -u github.com/golang/dep/cmd/dep
```

## Auth0

All request will the authenticated against **Auht0**.
If you wanna disable that, just remove `v1.Use(middlewares.Authentication)`.

You can configure you **Auth0** settings in `.env` file.

```
AUTH0_JWKS=https://{YOUR_AUTH0_TENANT_DOMAIN}/.well-known/jwks.json
AUTH0_AUDIENCE=http://{{YOUR_AUTH0_API_DOMAIN}
AUTH0_ISSUER=https://{YOUR_AUTH0_TENANT_DOMAIN}.auth0.com/
```
[Auth0](https://auth0.com/)

## Swagger

`swag` is a swagger docs generator.

### Installing

Install `swag` using `go get`

```sh
go get -u github.com/golang/dep/cmd/dep
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
docker run -p 9000:9000 your-image .
```