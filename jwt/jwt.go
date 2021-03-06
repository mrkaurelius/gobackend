package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var secret []byte = []byte("secret")

//CreateToken
func CreateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

// VerifyToken
func VerifyToken(tokenString string) (string, bool) {
	// sample token string taken from the New example

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	if err != nil {
		fmt.Println(err)
		//return "", false
	}

	if token == nil {
		return "", false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Print("claims: ")
		fmt.Println(claims["username"])

		// type assert
		s, _ := claims["username"].(string)
		return s, true
	}

	fmt.Println(err)

	return "", false
}

func main() {
	// auth user

}
