package common

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
)

var codeMsg = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := codeMsg[code]
	if !ok {
		return codeMsg[ERROR]
	}
	return msg
}
