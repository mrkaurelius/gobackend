package jwt_test

import (
	"testing"

	"github.com/mrkaurelius/gobackend/jwt"
)

// trivial test
func TestCreateToken(t *testing.T) {
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1yayJ9.M5v2V4WxhRxUSRyJ_JD2-KOm94zk1TvFyZHDcBcDHUs"

	tokenString := jwt.CreateToken("mrk")
	if expected != tokenString {
		t.Error("error")
		t.Log("token:", tokenString)
	}
	t.Log("expected: ", expected)
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1yayJ9.M5v2V4WxhRxUSRyJ_JD2-KOm94zk1TvFyZHDcBcDHUs"

	if jwt.VerifyToken(token) {

	} else {
		t.Error("error")
	}
}