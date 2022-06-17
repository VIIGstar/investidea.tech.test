package order_handlers

import (
	"github.com/gin-gonic/gin"
	"investidea.tech.test/internal/entities"
	query_params "investidea.tech.test/internal/query-params"
	"investidea.tech.test/pkg/auth"
	"investidea.tech.test/pkg/database"
	"net/http"
	"strconv"
)

func (h orderHandler) Accept(c *gin.Context) {
	session, err := auth.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "missing user from context")
		return
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		c.JSON(http.StatusBadRequest, "Empty ID")
		return
	}

	order, err := h.repo.Database().
		Order().
		Update(
			c,
			entities.Order{
				Status: entities.OrderStatusAccepted.String(),
			},
			query_params.OrderParams{
				CommonQueryParams: database.CommonQueryParams{
					ID: id,
				},
				Status:   entities.OrderStatusPending.String(),
				SellerID: session.UserID,
			},
			true)

	if err != nil {
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, order)
}
