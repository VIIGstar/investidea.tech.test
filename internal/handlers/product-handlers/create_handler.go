package product_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"investidea.tech.test/internal/entities"
	"investidea.tech.test/pkg/auth"
	"net/http"
)

func (h productHandler) Create(c *gin.Context) {
	session, err := auth.GetUserFromContext(c)
	if err != nil {
		logrus.Error("session empty")
		c.JSON(http.StatusBadRequest, "missing user from context")
		return
	}
	product := entities.Product{}
	if err = c.ShouldBindJSON(&product); err != nil {
		logrus.Error(fmt.Sprintf("Invalid request, err: %v", err))
		c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	product.SellerID = session.UserID
	err = h.repo.Database().Product().Create(&product)
	if err != nil {
		logrus.Error(fmt.Sprintf("Create product failed, err: %v", err))
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, product)
}
