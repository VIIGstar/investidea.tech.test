package order_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"investidea.tech.test/internal/dtos"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/pkg/auth"
	"net/http"
	"strings"
)

func (h orderHandler) Create(c *gin.Context) {
	session, err := auth.GetUserFromContext(c)
	if err != nil {
		logrus.Error("session empty")
		c.JSON(http.StatusBadRequest, "missing user from context")
		return
	}
	dto := dtos.OrderDTO{}
	if err = c.ShouldBindJSON(&dto); err != nil {
		logrus.Error(fmt.Sprintf("Invalid request, err: %v", err))
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if errs := dto.Validate(); errs != nil {
		c.JSON(http.StatusBadRequest, strings.Join(errs, "\n"))
		return
	}
	entity, _ := dto.ToEntity()
	order, _ := entity.(entities.Order)
	order.BuyerID = session.UserID
	err = h.repo.Database().Order().Create(&order)
	if err != nil {
		logrus.Error(fmt.Sprintf("Create order failed, err: %v", err))
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, order)
}
