package util

import (
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var Key string = ""

var PriKey = `-----BEGIN PRIVATE KEY-----
Proc-Type: Golang

MIICXAIBAAKBgQCl0n90KepaQsfZ0PtQUjQdWl1jFBAv4saDpFIIoNHSNoXYDJSA
n6tdF/y6prV+roZwiUHkj6JniL87O5htyRDaA8jyCLbDqCD3xA6ZhoQXJGBulHsQ
tPjMl/ti7i0sLAI1ITUG0N9uKz63y1GNleH6Z77yXPLXA163rJWp/59JaQIDAQAB
AoGAaLKOfKbWbCUMEDAN7Xnpo7rhieQqEEIL+yQuE3qO4D5CpdXTkZ6+0OmLTcJi
h9jj03wmWYyDCwHhVCHdrTutYiWhvw909A0H0YFDN69voGQYQzyPOhpcpNCBMapq
TLsOERaZando4/kg6V4ZR6sffL+QGlerg0S2a5chlHmT2QECQQDOhwX3XxaRmLFc
FhTclVXFtMlpUF6zMDLlIr1ATJlHyKX+Z4Tm1CCd1OggMoazBMUlVnLbB0LUpdO3
yZCnspphAkEAzYtK0nqtan/LwQlSV90sfP/5Zt2d/tiydo6EVdJsg1xYm7s4GgUY
ZWLRMLghhHqH4F/crLPGas88KCPF+f1cCQJAMg4vuADn62nisrr3W28mVsC2gDvm
d++apkaBL/BgxjEvajWU1I1dSmOrzwHv+7uQPLhzJfrgi8GVStojoUF8gQJBAMNT
dlQUfo2xSih0OksDBI75FcG6IrWWqeleP8XqxHIEpLPBM2wOoYNfZ5nbsQZbBpqj
nmqEIK9JpkoizLVnHjkCQD8mZQaBVp3xIT55Y7Aw+jFYPxEV9xuLmEihcGXuB3vE
XLl8MUNs4c80471BHagOkqd8ZrD/7PgLwP+hFHn6LGc=
-----END PRIVATE KEY-----
`

var PubKey = `-----BEGIN PUBLIC KEY-----
Proc-Type: Golang

MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCl0n90KepaQsfZ0PtQUjQdWl1j
FBAv4saDpFIIoNHSNoXYDJSAn6tdF/y6prV+roZwiUHkj6JniL87O5htyRDaA8jy
CLbDqCD3xA6ZhoQXJGBulHsQtPjMl/ti7i0sLAI1ITUG0N9uKz63y1GNleH6Z77y
XPLXA163rJWp/59JaQIDAQAB
-----END PUBLIC KEY-----
`

// 产生json web token
func GenToken(uid string, duration time.Duration) string {
	t := time.Now()
	claims := &jwt.StandardClaims{
		Id:        uid,                           //用户ID
		IssuedAt:  int64(t.Unix()),               //声明标识了JWT的时间
		NotBefore: int64(t.Unix()),               //不早于
		ExpiresAt: int64(t.Add(duration).Unix()), //到期时间
		Issuer:    "coffee377@dingtalk.com",      //发行方
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(Key))
	if err != nil {
		logs.Error(err)
		return ""
	}

	return ss
}

// secretFunc validates the secret format.
func secretFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(Key), nil
	}
}

func ValidToken(tokenString string) (valid bool, uid string) {

	token, _ := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, secretFunc())

	if token != nil {
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
			return token.Valid, claims.Id
		}
	}

	return false, ""
}
