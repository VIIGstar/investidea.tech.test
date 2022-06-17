package order_handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	query_params "investidea.tech.test/internal/query-params"
	"net/http"
)

func (h orderHandler) View(c *gin.Context) {
	var req = query_params.OrderParams{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		logrus.Error(fmt.Sprintf("Invalid request, err: %v", err))
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	orders, err := h.repo.Database().Order().Find(c, req, false)
	if err != nil {
		logrus.Error(fmt.Sprintf("get list order failed, err: %v", err))
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, orders)
}
