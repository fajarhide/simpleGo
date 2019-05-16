package controllers
 
import (
	"echo_with_mysql_mvc/models"
	"net/http"
	"github.com/labstack/echo"
)
 
func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	println("foo")
	return c.JSON(http.StatusOK,result)
}