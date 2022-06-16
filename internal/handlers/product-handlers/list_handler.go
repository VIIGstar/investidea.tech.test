package product_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	query_params "investidea.tech.test/internal/query-params"
	"net/http"
)

func (h productHandler) Search(c *gin.Context) {
	var req = query_params.ProductParams{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Error(fmt.Sprintf("Invalid request, err: %v", err))
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	products, err := h.repo.Database().Product().Find(c, req, false)
	if err != nil {
		logrus.Error(fmt.Sprintf("get list product failed, err: %v", err))
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, products)
}
