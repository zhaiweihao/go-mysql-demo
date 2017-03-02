package service

import (
    "helloGo/goDb/model"
    "strconv"
    "net/http"
    "github.com/labstack/echo"
)

func Create(c echo.Context) error {
    var user model.User
    id,_ :=strconv.ParseInt(c.FormValue("Id"),10,64)
    user.Id=id
    user.Name=c.FormValue("Name")
    err :=user.Create()
    if err !=nil{
        return echo.NewHTTPError(http.StatusInternalServerError)
    }
    return c.JSON(http.StatusOK,APIResult{Success:true,Result:map[string]string{"status":"ok"}})
}

func Retrive(c echo.Context) error {
    id,_ :=strconv.ParseInt(c.Param("id"),10,64)
    user,err :=model.User{}.Get(id)

    if err !=nil{
        return echo.NewHTTPError(http.StatusNotFound)
    }
    result :=make(map[string]interface{})
    if user !=nil{
        result["Name"]=user.Name
        result["Id"]=user.Id
    }
    return c.JSON(http.StatusOK,APIResult{Success:true,Result:result})
}