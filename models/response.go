package models

// Response 统一的响应码
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// RespCategory 响应获取分类数据
type RespCategory struct {
	Response
	Categories []Category `json:"category"`
}

// RespGame 响应游戏列表功能
type RespGame struct {
	Response
	Games []Game `json:"games"`
}
