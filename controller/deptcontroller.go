package controller

import (
	"github.com/labstack/echo"
	"testgopb/pb/pbdept"
	"net/http"
	"testGo/client"
	"golang.org/x/net/context"
)

func GetDepts(c echo.Context) error {
	deptRes, err := client.DeptClient.GetDepartments(context.TODO(), &pbdept.DepartmentsReq{})
	if err != nil {
		panic(err)
	}
	resResult := ResponseResult{
		Obj:deptRes,
	}

	return c.JSON(http.StatusOK, resResult)
}