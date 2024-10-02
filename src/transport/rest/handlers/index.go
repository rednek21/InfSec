package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func InitTemplates(router *gin.Engine) {
	templateDir := "web/templates/"

	router.LoadHTMLGlob(filepath.Join(templateDir, "*.html"))
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
