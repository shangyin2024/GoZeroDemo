// Code generated by goctl. DO NOT EDIT.
package types

type CommonPagination struct {
	Page     uint32 `json:"page,optional,default=1"`      // 页码
	PageSize uint32 `json:"pageSize,optional,default=10"` // 每页数量
}

type DemoListItem struct {
	Abbr string `json:"abbr"` // 缩写
	Name string `json:"name"` // 名称
}

type DemoListReq struct {
	CommonPagination
}

type DemoListRes struct {
	Total uint32          `json:"total"` // 总数
	Items []*DemoListItem `json:"items"`
}