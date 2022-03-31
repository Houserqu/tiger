package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) {
	r.GET("/admin", IndexView)
	r.GET("/editor", EditorView)
}

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func EditorView(c *gin.Context) {
	c.HTML(http.StatusOK, "amis.html", gin.H{})
}
