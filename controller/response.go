package controller

import (
	"context"
)

var (
	PageSize = 5
)

type ResponseResult struct {
	Boo bool
	Message string
	PageInfo *Page
	Obj interface{}
}
type Page struct {
	PageNum int
	TotalPage int
	TotalCount int64
	NaviPageNums []int
}

func GetTotalCount(index string, typee string) int64 {
	ctx := context.Background()
	/*cli, err :=es.NewClient()
	if err != nil {
		panic(err)
	}*/
	count, _ := cli.Count().Index(index).Type(typee).Do(ctx)
	return count
}

func GetTotalPage(pageSize int, totalCount int64) int{
	n := totalCount%int64(pageSize)
	tp := totalCount/int64(pageSize)
	if n == 0 {
		return int(tp)
	}else {
		return int(tp+1)
	}
}

func GetNaviPageNums(pageNum int, totalPage int, totalCount int64) []int {
	pSize := int64(PageSize)
	if totalCount > 0 {
		if (pSize * 4) < totalCount {
			if (pageNum - 2) <= 1 {
				return []int{1, 2, 3, 4, 5}
			} else if (pageNum + 2) >= totalPage {
				return []int{totalPage - 4, totalPage - 3, totalPage - 2, totalPage - 1, totalPage}
			} else {
				return []int{pageNum - 2, pageNum - 1, pageNum, pageNum + 1, pageNum + 2}
			}
		} else if (pSize * 3) < totalCount {
			return []int{1, 2, 3, 4}
		} else if (pSize * 2) < totalCount {
			return []int{1, 2, 3}
		} else if (pSize * 1) < totalCount {
			return []int{1, 2}
		} else if (pSize * 1) >= totalCount {
			return []int{1}
		}
	}
	return []int{}

}