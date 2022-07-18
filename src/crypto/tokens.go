package crypto

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"strconv"
	"time"
	"users_api/src/errorss"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const added_secret = "treeVerde"
const token_secret = "temporal"
const token_expire_minutes = 60 * 4

type TokenModel struct {
	IdUser uint
}

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
		Id:        id,
		ExpiresAt: time.Now().Add(time.Minute * token_expire_minutes).Unix(),
	})

	token, err := claims.SignedString([]byte(token_secret))
	if err != nil {
		panic(&errorss.ErrorResponseModel{HttpStatus: 500, Cause: "error in generating user access"})
	}

	return token
}

func ParseRequiredToken(plainToken string) *TokenModel {

	rawToken, err := jwt.Parse(plainToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_secret), nil
	})

	errMsg := ""
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errMsg = "token wrong format"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			errMsg = "Token expired"
		} else {
			errMsg = "Token invalid: A"
		}
	}

	if errMsg != "" {
		panic(errorss.ErrorResponseModel{HttpStatus: 401, Cause: errMsg})
	}

	if !rawToken.Valid {
		panic(errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Token invalid: B"})
	}

	token, errMsgGetInfo := getInfoFromClaims(rawToken)
	if errMsgGetInfo != "" {
		panic(errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Error getting info from token" + errMsgGetInfo})
	}

	return token
}

func ParseOptionalToken(plainToken string) *TokenModel {

	rawToken, _ := jwt.Parse(plainToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_secret), nil
	})

	token, _ := getInfoFromClaims(rawToken)

	return token
}

func getInfoFromClaims(rawToken *jwt.Token) (*TokenModel, string) {
	if claims, ok := rawToken.Claims.(jwt.MapClaims); !ok {
		return nil, ": a1"
	} else {
		if idUserSt, ok := claims["jti"].(string); !ok {
			return nil, ": a2"
		} else {
			if idUser, err := strconv.Atoi(idUserSt); err != nil {
				return nil, ": a3"
			} else {
				idUserU := uint(idUser)
				token := TokenModel{IdUser: idUserU}
				return &token, ""
			}
		}

	}

}
