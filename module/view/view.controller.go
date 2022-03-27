package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) {
	r.GET("/admin", IndexView)
}

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
