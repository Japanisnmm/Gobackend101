package gobackendauth

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/Japanisnmm/GoBackend101/config"
	"github.com/Japanisnmm/GoBackend101/modules/users"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
	Admin   TokenType = "admin"
	ApiKey  TokenType = "apikey"
)

type gobackendAuth struct {
	mapClaims *gobackendMapClaims
	cfg       config.IJwtConfig
}

type gobackendAdmin struct {
	*gobackendAuth
}





type gobackendMapClaims struct {

	Claims *users.UserClaims `json:"claims"`
	jwt.RegisteredClaims
}
type IGobackendAuth interface {
	SignToken() string
}
type IGobackendAdmin interface {
	SignToken() string
}



func jwtTimeDurationCal(t int) *jwt.NumericDate{
	return jwt.NewNumericDate(time.Now().Add(time.Duration(int64(t)*int64(math.Pow10(9)))))
}
func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func (a *gobackendAuth) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	ss, _ := token.SignedString(a.cfg.SecretKey())
	return ss
}
func (a *gobackendAdmin) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	ss, _ := token.SignedString(a.cfg.AdminKey())
	return ss
}

func ParseToken(cfg config.IJwtConfig, tokenString string) (*gobackendMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &gobackendMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method is invalid")
		}
		return cfg.SecretKey(), nil
	})
	if err != nil {
		if  errors.Is(err , jwt.ErrTokenMalformed) {
			return nil , fmt.Errorf("token format is in valid")
		} else if errors.Is(err, jwt.ErrTokenExpired){
			return nil,fmt.Errorf("token had expired")
		}else {
			return nil ,fmt.Errorf("parse token failed: %v", err)
		}
	}
	if claims, ok := token.Claims.(*gobackendMapClaims); ok {
		return claims,nil
	}else {
		return nil,fmt.Errorf("claims type is invalid")
	}

}


func ParseAdminToken(cfg config.IJwtConfig, tokenString string) (*gobackendMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &gobackendMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method is invalid")
		}
		return cfg.AdminKey(), nil
	})
	if err != nil {
		if  errors.Is(err , jwt.ErrTokenMalformed) {
			return nil , fmt.Errorf("token format is in valid")
		} else if errors.Is(err, jwt.ErrTokenExpired){
			return nil,fmt.Errorf("token had expired")
		}else {
			return nil ,fmt.Errorf("parse token failed: %v", err)
		}
	}
	if claims, ok := token.Claims.(*gobackendMapClaims); ok {
		return claims,nil
	}else {
		return nil,fmt.Errorf("claims type is invalid")
	}

}



func NewGobackendAuth (tokenType TokenType, cfg config.IJwtConfig, claims *users.UserClaims) (IGobackendAuth, error) {
	switch tokenType {
	case Access:
		return newAccessToken(cfg, claims), nil
	case Refresh :
		return newRefreshToken(cfg, claims),nil
	case Admin:
		return newAdminToken(cfg ),nil
	default:
		return nil, fmt.Errorf("unknown token type")
	}
}

func RepeatToken(cfg config.IJwtConfig, claims *users.UserClaims, exp int64) string{
	obj := &gobackendAuth{
		cfg: cfg,
		mapClaims: &gobackendMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "Gobackendshop-api",
				Subject:   "access-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeRepeatAdapter(exp),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},

		},
	}
	return obj.SignToken()
}






func newAccessToken(cfg config.IJwtConfig, claims *users.UserClaims) IGobackendAuth {
	return &gobackendAuth{
		cfg: cfg,
		mapClaims: &gobackendMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "Gobackendshop-api",
				Subject:   "access-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDurationCal(cfg.AccessExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}



func newRefreshToken(cfg config.IJwtConfig, claims *users.UserClaims) IGobackendAuth {
	return &gobackendAuth{
		cfg: cfg,
		mapClaims: &gobackendMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "Gobackendshop-api",
				Subject:   "refresh-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDurationCal(cfg.RefreshExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}


func newAdminToken(cfg config.IJwtConfig) IGobackendAuth {
	return &gobackendAdmin{
		&gobackendAuth{
			cfg: cfg,
			mapClaims: &gobackendMapClaims{
				Claims: nil,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "Gobackendshop-api",
					Subject:   "admin-token",
					Audience:  []string{ "admin"},
					ExpiresAt: jwtTimeDurationCal(300),
					NotBefore: jwt.NewNumericDate(time.Now()),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
				},
			},
		},


	}
}
