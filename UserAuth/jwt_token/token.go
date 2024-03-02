package jwt_token

import (
	db "UserAuth/database"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const LiveTimeOfToken = 1000

var SigningMethod = jwt.SigningMethodHS512

type Claims struct {
	Guid string `json:"guid"`
	jwt.StandardClaims
}

var secret = []byte("A&'/}Z57M(2hNg=;LE?")

func CreateRefreshToken(guid string, query func(string, string) error) (string, error) {

	var err error
	var tokenCrypt []byte

	token := make([]byte, 10)
	for i := range token {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		token[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	if tokenCrypt, err = bcrypt.GenerateFromPassword(token, 14); err == nil {
		if err = query(string(tokenCrypt), guid); err == nil {
			var tokenStr = base64.StdEncoding.EncodeToString(token)
			return tokenStr, err
		}
	}
	return "", err
}

func GetNewAccessToken(guid string) (string, error) {

	claims := Claims{
		guid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * LiveTimeOfToken).Unix(),
		},
	}

	access := jwt.NewWithClaims(SigningMethod, claims)

	return access.SignedString(secret)
}

func ParseVerifiedAccessToken(token string) (*Claims, error) {

	access, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if access.Valid {
		return access.Claims.(*Claims), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("that's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return access.Claims.(*Claims), fmt.Errorf("timing is everything")
		}
	}
	return nil, fmt.Errorf("couldn't handle this token")
}

func RefreshTokenValidate(guid, refresh string) error {
	var err error
	var dbRef *db.RefreshToken
	var decodeRef []byte
	if dbRef, err = db.ReadRefreshToken(guid); err == nil {
		if decodeRef, err = base64.StdEncoding.DecodeString(refresh); err == nil {
			if err = bcrypt.CompareHashAndPassword([]byte(dbRef.Refresh), decodeRef); err == nil {
				return nil
			}
		}
	}
	return err
}
