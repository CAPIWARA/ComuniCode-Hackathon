package users

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jwtCustomClaims struct {
	Id string
	jwt.StandardClaims
}

var exp = time.Now().Add(time.Hour * time.Duration(72)).Unix()
var iss = "capiwara"
var secretKey = "secretkey"

func Encode(uuid string) (string, error) {
	claims := &jwtCustomClaims{
		uuid,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Issuer:    iss,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwt.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return token, nil
}

func Decode(tokenString string) (*jwtCustomClaims, error) {
	log.Println("Decoding token: ", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Printf("problem to decode jwt: %v", err)
		return nil, err
	}
}

func (login *Login) Auth() (string, error) {
	fmt.Printf("login: %v", login)
	res, err := FindByEmail(login.Email)
	fmt.Printf("res: %v", res)
	fmt.Printf("errr: %v", err)

	return "", nil
}
