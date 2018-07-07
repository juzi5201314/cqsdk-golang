package cqsdk_golang

/*
#include <stdint.h>
#cgo LDFLAGS: -L ./ -lcqp
#include "cqp.h"
 */
import "C"
import "github.com/axgle/mahonia"

const (
	EVENT_IGNORE        = iota     //忽略事件
	EVENT_BLOCK                    //拦截事件
	REQUEST_ALLOW       = iota + 1 //请求_通过
	REQUEST_DENY                   //请求_拒绝
	REQUEST_GROUPADD    = iota + 1 //请求_群添加
	REQUEST_GROUPINVITE            //请求_群邀请
	CQLOG_DEBUG         = 0        //调试 灰色
	CQLOG_INFO          = 10       //信息 黑色
	CQLOG_INFOSUCCESS   = 11       //信息(成功) 紫色
	CQLOG_INFORECV      = 12       //信息(接收) 蓝色
	CQLOG_INFOSEND      = 13       //信息(发送) 绿色
	CQLOG_WARNING       = 20       //警告 橙色
	CQLOG_ERROR         = 30       //错误 红色
	CQLOG_FATAL         = 40       //致命错误 深红
)

var AUTHCODE C.int32_t

func SetAuthCode(c int32)  {
	AUTHCODE = C.int32_t(c)
}

func AddLog(p int32, h string, msg string) int32 {
	return int32(C.CQ_addLog(AUTHCODE, C.int32_t(p), gb18030(h), gb18030(msg)))
}

func SendPrivateMsg(from_qq int64, message string) int32 {
	return int32(C.CQ_sendPrivateMsg(AUTHCODE, C.int64_t(from_qq), gb18030(message)))
}

func SendGroupMsg(group_id int64, message string) int32 {
	return int32(C.CQ_sendGroupMsg(AUTHCODE, C.int64_t(group_id), gb18030(message)))
}

func SendDiscussMsg(discuss_id int64, message string) int32 {
	return int32(C.CQ_sendDiscussMsg(AUTHCODE, C.int64_t(discuss_id), gb18030(message)))
}

func DeleteMsg(message_id int64) int32 {
	return int32(C.CQ_deleteMsg(AUTHCODE, C.int64_t(message_id)))
}

func SendLike(qq int64) int32 {
	return int32(C.CQ_sendLike(AUTHCODE, qq))
}

func SetGroupKick(group_id int64, qq int64, rejectaddrequest bool) int32 {
	return int32(C.CQ_setGroupKick(AUTHCODE, C.int64_t(group_id), C.int64_t(qq), C.int32_t(boolToInt32(rejectaddrequest))))
}

func SetGroupBan(group_id int64, qq int64, duration int64) int32 {
	return int32(C.CQ_setGroupBan(AUTHCODE, C.int64_t(group_id), C.int64_t(qq), C.int64_t(duration)))
}

func SetGroupAdmin(group_id int64, qq int64, isadmin bool) int32 {
	return int32(C.CQ_setGroupAdmin(AUTHCODE, C.int64_t(group_id), C.int64_t(qq), C.int32_t(boolToInt32(isadmin))))
}

func SetGroupWholeBan(group_id int64, enableban bool) int32 {
	return int32(C.CQ_setGroupWholeBan(AUTHCODE, C.int64_t(group_id), C.int32_t(boolToInt32(enableban))))
}

func SetGroupAnonymousBan(group_id int64, anomymous string, duration int64) int32 {
	return int32(C.CQ_setGroupAnonymousBan(AUTHCODE, C.int64_t(group_id), gb18030(anomymous), C.int64_t(duration)))
}

func SetGroupAnonymous(group_id int64, enableanomymous bool) int32 {
	return int32(C.CQ_setGroupAnonymous(AUTHCODE, C.int64_t(group_id), C.int32_t(boolToInt32(enableanomymous))))
}

func SetGroupCard(group_id int64, qq int64, newcard string) int32 {
	return int32(C.CQ_setGroupCard(AUTHCODE, C.int64_t(group_id), C.int64_t(qq), gb18030(newcard)))
}

func SetGroupLeave(group_id int64, isdismiss bool) int32 {
	return int32(C.CQ_setGroupLeave(AUTHCODE, C.int64_t(group_id), C.int32_t(boolToInt32(isdismiss))))
}

func SetGroupSpecialTitle(group_id int64, qq int64, title string, duration int64) int32 {
	return int32(C.CQ_setGroupSpecialTitle(AUTHCODE, C.int64_t(group_id), C.int64_t(qq), gb18030(title), C.int64_t(duration)))
}

func SetDiscussLeave(discuss_id int64) int32 {
	return int32(C.CQ_setDiscussLeave(AUTHCODE, C.int64_t(discuss_id)))
}

func SetFriendAddRequest(responseflag string, responseoperation int32, remark string) int32 {
	return int32(C.CQ_setFriendAddRequest(AUTHCODE, gb18030(responseflag), C.int32_t(responseoperation), gb18030(remark)))
}

func SetGroupAddRequest(responseflag string, requesttype int32, responseoperation int32, reason string) int32 {
	return int32(C.CQ_setGroupAddRequestV2(AUTHCODE, gb18030(responseflag),  C.int32_t(requesttype),  C.int32_t(responseoperation), gb18030(reason)))
}

func GetGroupMemberInfo(group_id int64, qq int64 , oncache bool) string {
	return C.GoString(C.CQ_getGroupMemberInfoV2(AUTHCODE, C.int64_t(group_id), C.int64_t(qq), C.int32_t(boolToInt32(oncache))))
}

func GetStrangerInfo(qq int64 , oncache bool) string {
	return C.GoString(C.CQ_getStrangerInfo(AUTHCODE, C.int64_t(qq), C.int32_t(boolToInt32(oncache))))
}

func GetCookies() string {
	return C.GoString(C.CQ_getCookies(AUTHCODE))
}

func GetCsrfToken() int32 {
	return int32(C.CQ_getCsrfToken(AUTHCODE))
}

func GetLoginQQ() int64 {
	return int64(C.CQ_getLoginQQ(AUTHCODE))
}

func GetLoginNick() string {
	return C.GoString(C.GoString(C.CQ_getLoginNick(AUTHCODE)))
}

func GetAppDirectory() string {
	return C.GoString(C.GoString(C.CQ_getAppDirectory(AUTHCODE)))
}

func SetFatal(message string) int32 {
	return int32(C.CQ_setFatal(AUTHCODE, gb18030(message)))
}

func CQ_getRecord(file string, outformat string) string {
	return C.GoString(AUTHCODE, gb18030(file), gb18030(outformat))
}

func gb18030(c string) *C.char {
	return C.CString(mahonia.NewEncoder("gb18030").ConvertString(c))
}

func boolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}