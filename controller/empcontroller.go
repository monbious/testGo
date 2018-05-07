package controller

import (
	"github.com/labstack/echo"
	"net/http"

	"strconv"
	es "github.com/olivere/elastic"
	ecli "testGo/client"
	"golang.org/x/net/context"
	"testgopb/pb/pbemp"
	"strings"
	//"fmt"
)

var (
	EsUrl = "http://192.168.1.109:9200"
	EmpIndex = "data"
	EmpType = "employee"
)

var (
	cli,_  = es.NewClient(es.SetURL(EsUrl))
)

func CheckUser(c echo.Context) error {
	name := c.FormValue("empName")
//	fmt.Println("empName :" + name)
	ctx := context.Background()
	cli, _ := es.NewClient()
	q := es.NewTermQuery("EmpName",name)
	sr, _ := cli.Search().Index(EmpIndex).Type(EmpType).Query(q).Do(ctx)
	resResult := &ResponseResult{}
	if sr.Hits.TotalHits == 0 {
		resResult.Boo = true
	}else {
		resResult.Boo = false
	}

	return c.JSON(http.StatusOK, resResult)
}

func DelEmp(c echo.Context) error {
	eid := c.QueryParam("eid")

	ecli.EmpClient.DelEmp(context.TODO(), &pbemp.EmployeesReq{EmpName:eid})

	ctx := context.Background()
	//cli, _ := es.NewClient()
	if strings.Contains(eid, "-") {
		eids := strings.Split(eid, "-")
		for _, eid := range eids  {
			cli.Delete().Index(EmpIndex).Type(EmpType).Id(eid).Do(ctx)
		}

	}else {
		cli.Delete().Index(EmpIndex).Type(EmpType).Id(eid).Do(ctx)
	}
	resresult := &ResponseResult{
		Message:"删除成功",
	}

	return c.JSON(http.StatusOK, resresult)
}

func UpdateEmp(c echo.Context) error {
	eid := c.QueryParam("eid")
	email := c.FormValue("email")
	gender := c.FormValue("gender")
	did := c.FormValue("dId")
	ieid, _ := strconv.Atoi(eid)
	idid, _ := strconv.Atoi(did)
	empreq := &pbemp.EmployeesReq{
		EmpId:int32(ieid),
		Email:email,
		Gender:gender,
		DeptId:int32(idid),
	}
	//fmt.Printf("%+v", empreq)
	res, err := ecli.EmpClient.Update(context.TODO(), empreq)
	if err != nil {
		panic(err)
	}
	resResult := &ResponseResult{
		Boo:res.Boo,
	}

	return c.JSON(http.StatusOK, resResult)
}

func GetEmp(c echo.Context) error {
	eid := c.QueryParam("eid")
	ctx := context.Background()
	res, err2 := cli.Get().Index(EmpIndex).Type(EmpType).Id(eid).Do(ctx)
	if err2 != nil {
		panic(err2)
	}
	resResult :=&ResponseResult{
		Obj:res,
	}

	return c.JSON(http.StatusOK, resResult)

}

func SaveEmp(c echo.Context) error {
	name := c.FormValue("empName")
	gender := c.FormValue("gender")
	email := c.FormValue("email")
	did := c.FormValue("dId")
	deptid, _ :=strconv.Atoi(did)
	empRes, err := ecli.EmpClient.SaveEmp(context.TODO(), &pbemp.EmployeesReq{EmpName:name, Gender:gender, Email:email,DeptId:int32(deptid)})

	if err != nil {
		panic(err)
	}
	resResult := &ResponseResult{}
	if empRes.Boo {
		resResult.Boo=true
	}else {
		resResult.Boo=false
	}
	resResult.Obj = empRes


	return c.JSON(http.StatusOK, resResult)
}

func GetEmployees(c echo.Context) error{

	pageNum := c.QueryParam("pageNum")
	//log.Println(pageNum)
	//fmt.Println(c.RealIP())
	pn, _ := strconv.Atoi(pageNum)

	totalcount := GetTotalCount(EmpIndex, EmpType)
	totalpage := GetTotalPage(PageSize, totalcount)
	if pn > totalpage {
		pn = totalpage
	}
	fromNum := (pn-1)*5

	ctx := context.Background()
	//cli, err :=es.NewClient()
	//if err != nil {
	//	panic(err)
	//}
	result, e := cli.Search().Index(EmpIndex).Type(EmpType).From(fromNum).Size(PageSize).Sort("EmpId", true).Pretty(true).Do(ctx)
	if e != nil {
		panic(e)
	}
	//log.Println(len(result.Hits.Hits))

	navipagenums := GetNaviPageNums(pn,totalpage,totalcount)
	resResult :=&ResponseResult{
		PageInfo:&Page{
			PageNum:pn,
			TotalPage:totalpage,
			TotalCount:totalcount,
			NaviPageNums:navipagenums,
		},
		Obj:result.Hits.Hits,
	}

	return c.JSON(http.StatusOK, resResult)
}