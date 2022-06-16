package session_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	query_params "investidea.tech.test/internal/query-params"
	"investidea.tech.test/pkg/auth"
	info_log "investidea.tech.test/pkg/info-log"
	"net/http"
)

// @Summary  Validate user then get access token
// @Tags     Session
// @Param    wallet_address  query  string  true  "public key address to user wallet"
// @Accept   json
// @Produce  json
// @Success  200      {object}  auth.Authentication
// @Failure  400,500  {object}  object{error=string}
// @Router   /api/v1/sessions/login [post]
func (h *sessionHandler) Login(c *gin.Context) {
	user, err := h.repo.Database().User().Find(c, query_params.GetUserParams{
		Address: c.Query("wallet_address"),
	}, false)
	if err != nil {
		h.logger.Error("error find user", info_log.ErrorToLogFields("details", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, auth.Authentication{
		AccessToken: auth.New().GenerateAccessToken(user.Role, fmt.Sprintf("%v", user.ID), c),
		Success:     true,
	})
}
