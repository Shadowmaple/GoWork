package main

import (
	"time"
)

// TokenClaims means a claim segment in a JWT.
type TokenClaims struct {
	ID        uint32 `json:"id"`
	Role      uint32 `json:"role"`
	TeamID    uint32 `json:"team_id"`
	ExpiresAt int64  `json:"expires_at"` // 过期时间（时间戳，10位）
}

func (c TokenClaims) Valid() error {
	now := time.Now().Unix()

	if !c.VerifyExpiresAt(now, false) {
		return ErrTokenExpired
	}

	return nil
}

// Compares the exp claim against cmp.
// If required is false, this method will return true if the value matches or is unset
// Valid valids the token according to the ExpiresAt and current time.
func (c *TokenClaims) VerifyExpiresAt(now int64, required bool) bool {
	if c.ExpiresAt == 0 {
		return !required
	}
	return now <= c.ExpiresAt
}
