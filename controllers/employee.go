package controllers
 
import (
	"github.com/fajarhide/simpleGo/models"
	"net/http"
	"github.com/labstack/echo"
)
 
func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	println("foo")
	return c.JSON(http.StatusOK,result)
}