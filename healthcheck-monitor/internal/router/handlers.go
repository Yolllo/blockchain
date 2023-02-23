package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetServiceAlive(c *gin.Context) {
	resp, err := r.Core.GetServiceAlive()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
