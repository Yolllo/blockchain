package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetNodeStatus(c *gin.Context) {
	resp, err := r.Core.GetNodeStatus()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (r *Router) ExecNode(c *gin.Context) {
	err := r.Core.ExecNode()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (r *Router) GetServiceAlive(c *gin.Context) {
	resp, err := r.Core.GetServiceAlive()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
