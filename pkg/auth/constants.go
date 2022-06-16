package auth

import "time"

const AccessTokenExpiry = time.Hour

const (
	AccessTokenKey     = "access-token"
	IssuerIdClaimKey   = "iss"
	IssuerRoleClaimKey = "sub"
)
