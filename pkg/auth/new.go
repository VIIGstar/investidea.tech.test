package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	info_log "investidea.tech.test/pkg/info-log"
	"logur.dev/logur"
	"net/http"
	"strings"
	"time"
)

const (
	BuyerRole  = RoleType("buyer")
	SellerRole = RoleType("seller")
)

type RoleType string

type Authenticator interface {
	GenerateAccessToken(role, id string, c *gin.Context) string
}

type impl struct {
}

func (i impl) GenerateAccessToken(role, id string, c *gin.Context) string {
	return newJWTService().generateToken(role, id, c, time.Now().Add(AccessTokenExpiry))
}

func New() Authenticator {
	return impl{}
}
func Authorize(role RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetUserFromContext(c)
		if err != nil {
			logrus.Info("[Auth] no user from context")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.Role != role.String() {
			logrus.Info("[Auth] permission denied")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// Middleware checks if request comes with valid access token
func Middleware(logger logur.LoggerFacade) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !skipTokenCheck(c.Request.URL.Path) {
			at := extractAccessToken(c)
			if at == "" {
				logger.Info("[Auth] no access token, proceed to extract refresh token")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// if access token present
			t, err := newJWTService().validateToken(at, c)
			if err != nil {
				logger.Error("[Auth] error when verify access token",
					info_log.ErrorToLogFields("details: ", err))
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Extract token metadata and store its tokenDetails into the same request context
			userDetails, err := extractTokenMetadata(t)
			if err != nil || userDetails == nil {
				logger.Error("[Auth] error when extract token metadata",
					info_log.ErrorToLogFields("details: ", err))
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			setUserContext(c, *userDetails)
			c.Next()
		}
	}
}

func extractAccessToken(c *gin.Context) string {
	atCookie, err := c.Request.Cookie(AccessTokenKey)
	if err != nil || len(atCookie.Value) == 0 {
		at := c.GetHeader(strings.ToUpper(AccessTokenKey))
		if len(at) == 0 {
			return ""
		}

		return at
	}
	return atCookie.Value
}

// extractTokenMetadata extracts metadata from token
func extractTokenMetadata(token *jwt.Token) (*UserDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrMapClaimNotFound
	}

	userID, ok := claims[IssuerIdClaimKey]
	if !ok {
		return nil, ErrUserIDNotFound
	}

	role, ok := claims[IssuerRoleClaimKey]
	if !ok {
		return nil, ErrUserAddressNotFound
	}

	return &UserDetails{
		UserID: cast.ToInt64(userID),
		Role:   role.(string),
	}, nil
}

func skipTokenCheck(uri string) bool {
	skipPath := []string{
		"/api/v1/liveness",
		"/api/v1/readiness",
		"/api/v1/sessions/login",
		"/api/v1/users/signup",
		"/api/v1/debug/pprof",
		"/swagger",
	}
	for _, s := range skipPath {
		if strings.Contains(uri, s) {
			return true
		}
	}
	return false
}

func (t RoleType) String() string {
	return string(t)
}
