package serializer

// Response 通用响应
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

// DataList 通用分页列表响应
type DataList[T interface{}] struct {
	List  []T  `json:"list"`
	Total uint `json:"total"`
}
