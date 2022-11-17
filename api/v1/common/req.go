package common

type PageReq struct {
	DateRange []string `p:"dateRange"` //日期范围
	PageNum   int      `p:"pageNum"`
	PageSize  int      `p:"pageSize"`
	OrderBy   string   //排序方式
}

type ListResp struct {
	CurrentPage int `json:"currentPage"`
	Total       int `json:"total"`
}
