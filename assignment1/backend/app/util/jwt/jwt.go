package jwt

import (
	"app/config/log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type ClaimData struct {
	UserID string `json:"user_id"`
}

type Claims struct {
	Data ClaimData
	jwt.RegisteredClaims
}

func CreateToken(claims ClaimData) (*string, error) {

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	now := time.Now().In(loc)

	duration, err := strconv.Atoi(viper.GetString("JWT_DURATION"))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		claims,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(duration) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	})
	secret := []byte(viper.GetString("JWT_SECRET"))

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &tokenString, nil
}

func Verify(token string) (*ClaimData, error) {

	secret := []byte(viper.GetString("JWT_SECRET"))
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	claims, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, err
	}

	return &claims.Data, nil
}
