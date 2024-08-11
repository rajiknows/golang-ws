package handlers

import (
	"time"
)

type Claim struct {
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expires"`
}
