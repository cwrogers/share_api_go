package mw

import (
	"net/http"
	"share/share-api/common/config"
	"share/share-api/models/app"
	"strings"
	"time"

	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthenticationResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

var jwtSecret = config.ApplicationConfig.JwtSecret

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appG := app.Gin{Ctx: ctx}
		token := strings.Trim(ctx.Request.Header.Get("Authorization"), " ")
		if token == "" {
			appG.Response(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			appG.Response(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.ApplicationConfig.Name,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", "", err
	}
	rtClaims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: nowTime.Add(time.Hour * 24 * 7).Unix(),
			Issuer:    config.ApplicationConfig.Name,
		},
	}

	//gen refresh token
	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := refreshTokenClaims.SignedString([]byte(jwtSecret))

	return token, refreshToken, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
