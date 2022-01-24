package crypto

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"users_api/src/errorss"

	"golang.org/x/crypto/bcrypt"
)

const added_secret = "treeVerde"

func GenerateToken() (plainText string, hash string) {

	//generamos Bites aleatorios
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(
			errorss.ErrorResponseModel{
				Cause:      "Error al generar clave de activacion",
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
