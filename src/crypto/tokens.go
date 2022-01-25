package crypto

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"time"
	"users_api/src/errorss"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const added_secret = "treeVerde"
const token_secret = "temporal"

func GenerateRandomHash() (plainText string, hash string) {

	//generamos Bites aleatorios
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(
			errorss.ErrorResponseModel{
				Cause:      "Error in generate activation code",
				HttpStatus: 500,
			},
		)
	}
	//De los bites aleatorios generamos texto-base64
	plainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	//Del token generamos su Hash
	hash = GetHash(plainText)

	return plainText, hash
}

func PasswordMatches(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+added_secret))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false
		default:
			panic(
				errorss.ErrorResponseModel{
					Cause:      "Error al comparar contraseña",
					HttpStatus: 500,
				},
			)
		}
	}

	return true
}

func GetHash(plainText string) string {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(plainText+added_secret), 12)
	hash := string(hashBytes[:])
	if err != nil {
		panic(
			errorss.ErrorResponseModel{
				Cause:      "Error al cifrar contraseña",
				HttpStatus: 500,
			},
		)
	}
	return hash
}

func GenerateToken(id string) string {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    id,
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(token_secret))
	if err != nil {
		panic(&errorss.ErrorResponseModel{HttpStatus: 500, Cause: "error in generating user access"})
	}

	return token
}

func ParseToken(plainToken string) interface{} {
	invalidToken := &errorss.ErrorResponseModel{HttpStatus: 401, Cause: "invalid token"}

	token, err := jwt.Parse(plainToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_secret), nil
	})

	if err != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "error validating token"})
	}

	if !token.Valid {
		panic(invalidToken)
	} else {
		return token.Claims
	}
}
