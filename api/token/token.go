package token

import (
	"errors"
	"log"
	"time"

	"github.com/dilshodforever/restaurant-auth/config"
	pb "github.com/dilshodforever/restaurant-auth/genprotos"
	"github.com/form3tech-oss/jwt-go"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

var tokenKey = config.Load().TokenKey

func GenereteJWTToken(user *pb.User) *Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["username"] = user.UserName
	claims["email"] = user.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting access token : ", err)
	}

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["user_id"] = user.Id
	rftclaims["username"] = user.UserName
	rftclaims["email"] = user.Email
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting refresh token : ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}



func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
