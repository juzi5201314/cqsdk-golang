package main

import "C"
import (
	"github.com/juzi5201314/cqsdk-golang"
)

//export AppInfo
func AppInfo() *C.char {
	return C.CString("9,cn.you.test")
}

//export Initialize
func Initialize(autocode int32) int32 {
	cqsdk_golang.SetAuthCode(autocode)
	return cqsdk_golang.EVENT_IGNORE
}

//此处注释与函数名字必须跟appid.json里所填写的一模一样
//export OnEnable
func OnEnable() int32 {
	cqsdk_golang.AddLog(cqsdk_golang.CQLOG_INFO, "启动", "test插件成功启动")
	return cqsdk_golang.EVENT_IGNORE
}

//export OnPrivateMessage
func OnPrivateMessage(sub_type int32, message_id int32, from_qq int64, msg *C.char, font int32) int32 {
	cqsdk_golang.SendPrivateMsg(from_qq, "收到了消息: " + cqsdk_golang.Utf8(C.GoString(msg)))
	return cqsdk_golang.EVENT_IGNORE
}

func main()  {
	
}
