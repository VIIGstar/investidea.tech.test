package user_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"investidea.tech.test/internal/dtos"
	"investidea.tech.test/internal/entities"
	database "investidea.tech.test/internal/services/database/mysql"
	"investidea.tech.test/pkg/auth"
	app_http "investidea.tech.test/pkg/http"
	"net/http"
	"strings"
)

// @Summary  Signup create new user
// @Tags     Investor
// @Param    data body dtos.UserDTO true "The input struct"
// @Accept   json
// @Produce  json
// @Success  200      {object}  auth.Authentication
// @Failure  400,500  {object}  object{error=string}
// @Router   /api/v1/users/signup [post]
func (s *userHandler) Signup(c *gin.Context) {
	user := dtos.UserDTO{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Invalid request, err: %v", err))
		c.JSON(http.StatusBadRequest, auth.Authentication{
			Error: app_http.HTTPBadRequestError,
		})
		return
	}

	if errs := user.Validate(); len(errs) > 0 {
		c.JSON(http.StatusBadRequest, auth.Authentication{
			Error: strings.Join(errs, "\n"),
		})
		return
	}

	iEntity, err := user.ToEntity()
	if err != nil {
		s.logger.Error(fmt.Sprintf("Parse to entity failed, err: %v", err))
		c.JSON(http.StatusInternalServerError, auth.Authentication{
			Error: app_http.HTTPInternalServerError,
		})
		return
	}

	entity, _ := iEntity.(entities.User)
	err = s.repo.Database().User().Create(&entity)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Insert into database failed, err: %v", err))
		if database.IsDuplicateErr(err) {
			c.JSON(http.StatusInternalServerError, auth.Authentication{
				Error: auth.ErrUserAlreadyRegistered.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, app_http.HTTPInternalServerError)
		return
	}

	c.JSON(http.StatusOK, auth.Authentication{
		AccessToken: auth.New().GenerateAccessToken(entity.Role, fmt.Sprintf("%v", entity.ID), c),
		Success:     true,
	})
}
