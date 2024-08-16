package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	//app.GET("/404-page", handlers.ErrorPage)
	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
}
