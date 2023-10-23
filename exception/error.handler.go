package exception

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/randijulio13/gogin/response"
)

func ErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			switch e := err.(type) {
			case NotFoundException:
				c.JSON(http.StatusNotFound, response.Web{
					Message: e.Error,
				})
			case ValidationException:
				c.JSON(http.StatusBadRequest, response.Web{
					Message: e.Error,
				})
			default:
				c.JSON(http.StatusInternalServerError, response.Web{
					Message: "Internal Server Error",
				})
			}
			c.Abort()
			return
		}
	}()
	c.Next()
}
