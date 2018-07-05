package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenError(obj interface{}, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, obj)
}
func BadRequest(obj interface{}, c *gin.Context) {
	c.JSON(http.StatusBadRequest, obj)
}

func Json(obj interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, obj)
}
