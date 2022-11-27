package helpers

import "github.com/gin-gonic/gin"

func SendResponse(c *gin.Context, status int64, msg string, data interface{}) {
	c.JSON(int(status), map[string]interface{}{"msg": msg, "data": data})
}
