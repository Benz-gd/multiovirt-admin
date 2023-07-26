package controllers


const(
	CodeNoRoute = 404
	CodeInvalidParam = 4000 +iota
	CodeUserExist
	CodeUserNotExist
	CodeRegisterFail
	CodeInvalidPassword
	CodeServerBusy
	CodeAuthNotExist
	CodeAuthFormatErr
	CodeAuthInvalidToken
	CodeUserNotLogin
	CodeOvirtConfExist
	CodeDataOperationErr
	CodeInsertDBErr
)
var ErrCodeMsgMap = map[int]string{
	CodeNoRoute: "网页找不到",
	CodeInvalidParam: "请求参数错误",
	CodeUserExist: "用户已存在",
	CodeUserNotExist: "用户不存在",
	CodeRegisterFail: "注册失败",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy: "服务繁忙",
	CodeAuthNotExist: "请求头Auth为空",
	CodeAuthFormatErr: "请求头中Auth格式有误",
	CodeAuthInvalidToken: "无效Token",
	CodeUserNotLogin: "用户未登录",
	CodeOvirtConfExist:"数据库配置已存在",
	CodeDataOperationErr:"数据操作错误",
	CodeInsertDBErr:"插入数据错误",
}

const(
	CodeLoginSuccess = 2000 + iota
	CodeResponseSuccess
	CodeRegisterSuccess
	CodeDataOperationSuccess
)
var SucCodeMsgMap = map[int]string{
	CodeLoginSuccess: "登录成功",
	CodeRegisterSuccess: "注册成功",
	CodeResponseSuccess: "Success",
	CodeDataOperationSuccess: "数据操作成功",
}


func GetErrMsg(code int) string{
	msg,ok := ErrCodeMsgMap[code]
	if ! ok {
		msg = ErrCodeMsgMap[CodeServerBusy]
	}
	return msg
}


func GetSucMsg(code int) string{
	return  SucCodeMsgMap[code]
}