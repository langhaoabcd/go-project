package models

//查询分页统一请求
type QueryRequest struct {
	PageIndex int	`json:"pageindex"`
	PageSize int	`json:"pagesize"`
	SortColumn string	`json:"sortcolumn"`
	SortDirection string	`json:"sortdirection"`
}