package middlewares

import (
	"os"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

//Authentication ...
func Authentication(context *gin.Context) {
	jwks := os.Getenv("AUTH0_JWKS")
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwks})
	audience := []string{os.Getenv("AUTH0_AUDIENCE")}

	configuration := auth0.NewConfiguration(client, audience, os.Getenv("AUTH0_ISSUER"), jose.RS256)
	validator := auth0.NewValidator(configuration)

	_, err := validator.ValidateRequest(context.Request)

	if err != nil {
		context.AbortWithStatus(401)
	} else {
		context.Next()
	}
}
