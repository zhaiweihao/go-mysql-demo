package main

import(
    "helloGo/goDb/service"
    "net/http"
    "github.com/labstack/echo"
)

func RegisterApi()  {
    e :=echo.New()
    e.GET("/",func (c echo.Context) error {
        return c.String(http.StatusOK,"hello db")
    })
    api :=e.Group("/api")
    v1 :=api.Group("/v1")

    user :=v1.Group("/customer")
    user.GET("/:id",service.Retrive)
    user.POST("",service.Create)
}