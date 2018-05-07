package client

import (
	"testgopb/pb"
	"testgopb/pb/pbemp"
	"testgopb/pb/pbdept"
	"testGo/conf"
)

var (
	EmpClient pbemp.EmployeesClient
	DeptClient pbdept.DepartmentsClient
)

func StartClient() {
	service := pb.NewService(&conf.Config.ConfigServiceMicro)
	service.Init()

	EmpClient = pbemp.NewEmployeesClient(pb.ResolveSVCName, service.Client())
	DeptClient = pbdept.NewDepartmentsClient(pb.ResolveSVCName, service.Client())
}