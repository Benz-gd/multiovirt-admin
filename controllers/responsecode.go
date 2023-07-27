package controllers

const (
	CodeNoRoute      = 404
	CodeServerBusy   = 500
	CodeInvalidParam = 4000 + iota
	CodeUserErr
	CodeInvalidPassword
	CodeUserNotLogin
	CodeRegisterFail
	CodeAuthInvalidToken
	CodeDataOperationErr
)

var ErrCodeMsgMap = map[int]string{
	CodeNoRoute:          "网页找不到",
	CodeServerBusy:       "服务繁忙",
	CodeInvalidParam:     "请求参数错误",
	CodeUserErr:          "用户信息错误",
	CodeInvalidPassword:  "用户名或密码错误",
	CodeUserNotLogin:     "用户未登录",
	CodeRegisterFail:     "注册失败",
	CodeAuthInvalidToken: "无效Token",
	CodeDataOperationErr: "数据操作错误",
}

const (
	CodeLoginSuccess = 2000 + iota
	CodeRegisterSuccess
	CodeResponseSuccess
	CodeDataOperationSuccess
)

var SucCodeMsgMap = map[int]string{
	CodeLoginSuccess:         "登录成功",
	CodeRegisterSuccess:      "注册成功",
	CodeResponseSuccess:      "Success",
	CodeDataOperationSuccess: "数据操作成功",
}

func GetErrMsg(code int) string {
	msg, ok := ErrCodeMsgMap[code]
	if !ok {
		msg = ErrCodeMsgMap[CodeServerBusy]
	}
	return msg
}

func GetSucMsg(code int) string {
	return SucCodeMsgMap[code]
}
