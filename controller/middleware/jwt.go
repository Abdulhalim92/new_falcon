package middleware

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"falcon/config"
	"falcon/pkg/logging"
	"falcon/repository/identity"
	"falcon/service"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func JwtMiddleware(cfg config.KeycloakParams, logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRetrospect := identity.NewIdentityManager(cfg, logger)
		successHandler(c, tokenRetrospect)
	}
}

func successHandler(c *gin.Context, tokenRetrospect service.Identity) {
	userToken := c.GetHeader("user")
	if userToken == "" {
		log.Println("unable to get token")
		c.JSON(http.StatusUnauthorized, "cannot get token")
		c.Abort()
		return
	}

	base64Str := viper.GetString("KeyCloak.RealmRS256PublicKey")
	publicKey, err := parseKeycloakRSAPublicKey(base64Str)
	if err != nil {
		log.Println(err, "unable to get public key")
		c.JSON(http.StatusInternalServerError, "unable to get public key")
		c.Abort()
		return
	}

	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		log.Println(err, "unable to parse token")
		c.JSON(http.StatusInternalServerError, "unable to parse token")
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	var ctx = c.Request.Context()

	c.Set("key_claims", claims)

	rptResult, appError := tokenRetrospect.RetrospectToken(ctx, token.Raw)
	if appError.Err != nil {
		panic(err)
	}
	if !*rptResult.Active {
		log.Println("token is not active")
		c.JSON(http.StatusUnauthorized, "token is not active")
		c.Abort()
		return
	}

	c.Next()
}

func parseKeycloakRSAPublicKey(base64Str string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	parsedKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, err
	}
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if ok {
		return publicKey, nil
	}
	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}
