package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey string = "jwttest"

	ErrTokenInvalid = errors.New("The token is invalid.")
	ErrTokenExpired = errors.New("The token is expired.")
)

// getJwtKey get jwtKey.
func getJwtKey() string {
	// if jwtKey == "" {
	// 	jwtKey = viper.GetString("jwt_secret")
	// }
	return jwtKey
}

// TokenPayload is a required payload when generates token.
type TokenPayload struct {
	ID      uint32        `json:"id"`
	Role    uint32        `json:"role"`
	TeamID  uint32        `json:"team_id"`
	Expired time.Duration `json:"expired"` // 有效时间
}

// TokenResolve means returned payload when resolves token.
type TokenResolve struct {
	ID     uint32 `json:"id"`
	Role   uint32 `json:"role"`
	TeamID uint32 `json:"team_id"`
	// ExpiresAt int64  `json:"expires_at"` // 过期时间（时间戳，10位）
}

// GenerateToken generates token.
func GenerateToken(payload *TokenPayload) (string, error) {
	claims := &TokenClaims{
		ID:        payload.ID,
		Role:      payload.Role,
		TeamID:    payload.TeamID,
		ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
	}
	fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getJwtKey()))
}

// ResolveToken resolves token.
func ResolveToken(tokenStr string) (*TokenResolve, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})
	fmt.Println(claims)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	t := &TokenResolve{
		ID:     claims.ID,
		Role:   claims.Role,
		TeamID: claims.TeamID,
	}
	return t, nil
}
