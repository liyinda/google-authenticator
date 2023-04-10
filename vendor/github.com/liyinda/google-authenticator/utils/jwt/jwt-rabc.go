package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	AcceptTokenKey      = "authorization"
	RefreshTokenKey     = "RFToken"
	TokenIssuer         = "testIssuer"
	mySecret            = "666test"
	TokenExpireDuration = time.Hour * 24 * 7
	//TokenExpireDuration        = time.Minute * 1
	RefreshTokenExpireDuration = time.Hour * 24 * 30
)

var (
	// 预设错误信息
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type MyClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenAccessToken(userName, role string) (aToken string, err error) {
	// 创建一个我们自己的声明的数据
	cl := MyClaims{
		UserName: userName,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    TokenIssuer,
		},
	}
	// 使用指定的签名方法创建签名对象
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, cl).SignedString([]byte(mySecret))
	return
}

func GenRefreshToken() (rToken string, err error) {
	// refresh token
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
		Issuer:    TokenIssuer,
	}).SignedString([]byte(mySecret))
	return
}

//it is used for ldap auth
//func ParseToken(tokenString string) (*MyClaims, error) {
//	// 解析token
//	var mc = new(MyClaims)
//	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
//		return []byte(mySecret), nil
//	})
//	if err != nil {
//		// 如果强转*jwt.ValidationError成功，对错误进行判断
//		if validationError, ok := err.(*jwt.ValidationError); ok {
//			/*
//				当validationError中的错误信息由错误的token结构引起时，
//				**************************************************
//				源码vErr.Errors |= ValidationErrorExpired，
//				或运算，只有都为0才为0，0000 0000|0000 0101 = 0000 0101
//				由于vErr.Errors的初始值为0，所以等价于将ValidationErrorMalformed赋值给validationError的Errors，
//				*****************************************************
//				如果没有赋值，Errors的初始值为0，那么validationError.Errors&jwt.ValidationErrorMalformed = 0，
//				赋值后造成validationError.Errors不为0，那么validationError.Errors&jwt.ValidationErrorMalformed != 0
//			*/
//			if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
//				return nil, TokenMalformed
//				// 以下与上方原理相同
//			} else if validationError.Errors&jwt.ValidationErrorExpired != 0 {
//				return nil, TokenExpired
//			} else if validationError.Errors&jwt.ValidationErrorNotValidYet != 0 {
//				return nil, TokenNotValidYet
//			} else {
//				return nil, TokenInvalid
//			}
//		}
//	}
//	if token != nil {
//		// 强转成jwtClaims
//		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
//			// 如果合法返回claims
//			return claims, nil
//		}
//		return nil, TokenInvalid
//	} else {
//		return nil, TokenInvalid
//	}
//}

func RefreshToken(aToken, rToken string) (newAToken string, err error) {
	// refresh token 无效直接返回
	if _, err = jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySecret), nil
	}); err != nil {
		return
	}

	// 从旧的token解析出claims数据
	var claims = new(MyClaims)
	_, err = jwt.ParseWithClaims(aToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySecret), nil
	})
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return GenAccessToken(claims.UserName, claims.Role)
	}
	return
}
