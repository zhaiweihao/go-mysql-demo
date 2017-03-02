package main

import (
    "helloGo/goDb/model"
    "github.com/labstack/echo"
)
var (
    e=echo.New()
)
func main()  {
    RegisterApi()
    model.InitDB("mysql","root:zhai123@tcp(127.0.0.1:3306)/mysqltest?charset=utf8&parseTime=True&loc=UTC")
    e.Start(":1114")
}