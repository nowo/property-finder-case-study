package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"property-finder-go-bootcamp-homework/internal/.config"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return _config.JWT_SECRETKEY, nil
}

type JWT struct {
	DumpClaim jwt.MapClaims
	Token     string
}

func New() *JWT {
	return &JWT{
		DumpClaim: jwt.MapClaims{
			"exp": _config.TOKEN_EXPIRATION_TIME,
		},
		Token: "",
	}
}

func (_jwt *JWT) SetUserId(userId uint) *JWT {
	_jwt.DumpClaim["userId"] = userId
	return _jwt
}

func (_jwt *JWT) GetUserId() float64 {
	userId := _jwt.DumpClaim["userId"]
	return userId.(float64)
}

func (_jwt *JWT) GetToken() string {
	token := _jwt.DumpClaim["token"]
	return token.(string)
}

func (_jwt *JWT) SetToken(token string) *JWT {
	_jwt.DumpClaim["token"] = token
	return _jwt
}

func (_jwt *JWT) SetIsVerified(isVerified bool) *JWT {
	_jwt.DumpClaim["is_verified"] = isVerified
	return _jwt
}

func (_jwt *JWT) CreateToken() *JWT {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, _jwt.DumpClaim)
	tokenString, _ := token.SignedString([]byte(_config.JWT_SECRETKEY))
	_jwt.SetToken(tokenString)
	return _jwt
}

func (_jwt *JWT) DecodeToken() *JWT {
	jwt.ParseWithClaims(_jwt.GetToken(), _jwt.DumpClaim, keyFunc)
	return _jwt
}
