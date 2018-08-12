package vo

// ResponseData	通用返回模型
type ResponseData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}
