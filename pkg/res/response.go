package res

type Res struct {
	Code    int    `json:"code"`    // 业务状态码
	Msg     string `json:"msg"`     // 业务状态码描述
	Success bool   `json:"success"` // 请求是否成功
}

type ResData struct {
	Res
	Data interface{} `json:"data"` // 响应数据
}

type ResPageData struct {
	ResData
	Page     int `json:"page"`     // 页数
	PageSize int `json:"pageSize"` // 分页大小
	Total    int `json:"total"`    // 数据总数
}
