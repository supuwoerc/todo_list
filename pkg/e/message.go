package e

var codeMessageMapping = map[int]string{
	SUCCESS:       "操作成功",
	ERROR:         "操作失败",
	InvalidParams: "请求参数错误",

	ErrorExistUser:    "用户已存在",
	ErrorNotExistUser: "用户不存在",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorNotCompare:            "密码不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
}

func GetMessage(code int) string {
	msg, ok := codeMessageMapping[code]
	if ok {
		return msg
	}
	return codeMessageMapping[ERROR]
}
