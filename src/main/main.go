package main

import (
	"net/http"
  "fmt"
	"github.com/labstack/echo"
)

func yallo(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func getUsuarios(c echo.Context) error{
  userName := c.QueryParam("name")
  aisleName := c.QueryParam("location")

  dataType := c.Param("data")
  if dataType == "string"{
    return c.String(http.StatusOK, fmt.Sprintf("El nombre del usuario es: %s\nEl pasillo en que esta es %s\n",userName, aisleName))
  }
  if dataType == "json" {
    return c.JSON(http.StatusOK, map[string]string{
      "name":userName,
      "aisle":aisleName,
    })
  }
  return c.JSON(http.StatusBadRequest, map[string]string{
    "Error": "no type",
  })


}

func main() {
  fmt.Println("hello")
	e := echo.New()
	e.GET("/", yallo)
  e.GET("/user/:data", getUsuarios)
	e.Logger.Fatal(e.Start(":8000"))
}
