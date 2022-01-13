package middleware

import (
	"crypto/rand"
	"fmt"
	"hitszedu-go/config"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var key []byte

func init() {
	_randInt, _ := rand.Int(rand.Reader, big.NewInt(999999))
	randInt := _randInt.String()
	key = []byte(config.GetString("jwt.key") + randInt)
}

type Claims struct {
	jwt.StandardClaims
	UserID string `json:"userID"`
}

func GenerateToken(userID string) (token string, success bool) {
	nowTime := time.Now()                    //当前时间
	expireTime := nowTime.Add(3 * time.Hour) //有效时间
	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.GetString("jwt.issuer"),
		},
		userID,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(key)
	if err != nil {
		return "", false
	} else {
		return token, true
	}
}

func ParseToken(signedToken string) (userID string, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return key, nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)

	if ok && token.Valid {
		userID = claims.UserID
		success = true
	} else {
		userID = ""
		success = false
	}
	return
}

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"state": "authNotExist",
				"data":  "",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"state": "authIllegal",
				"data":  "",
			})
			c.Abort()
			return
		}
		userName, ok := ParseToken(authHeader)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"state": "authIllegal",
				"data":  "",
			})
			c.Abort()
			return
		}
		c.Set("username", userName)
		c.Next()
	}
}
