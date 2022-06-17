package session_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"investidea.tech.test/internal/dtos"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/pkg/auth"
	info_log "investidea.tech.test/pkg/info-log"
	"net/http"
)

// @Summary  Validate user then get access token
// @Tags     Session
// @Param    data body  string  true  "public key address to user wallet"
// @Accept   json
// @Produce  json
// @Success  200      {object}  auth.Authentication
// @Failure  400,500  {object}  object{error=string}
// @Router   /api/v1/sessions/login [post]
func (h *sessionHandler) Login(c *gin.Context) {
	userDto := dtos.UserDTO{}
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	entityReq, _ := userDto.ToEntity()
	userReq := entityReq.(entities.User)

	user, err := h.repo.Database().User().Login(c, userReq.Email, userReq.Password)
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
