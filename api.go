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

var AUTOCODE C.int32_t

func SetAutoCode(c int32)  {
	AUTOCODE = C.int32_t(c)
}

func AddLog(p int32, h string, msg string) int32 {
	return int32(C.CQ_addLog(AUTOCODE, C.int32_t(p), gbk(h), gbk(msg)))
}

func SendPrivateMsg(from_qq int64, message string) int32 {
	return int32(C.CQ_sendPrivateMsg(AUTOCODE, C.int64_t(from_qq), gbk(message)))
}

func SendGroupMsg(group_id int64, message string) int32 {
	return int32(C.CQ_sendGroupMsg(AUTOCODE, C.int64_t(group_id), gbk(message)))
}

func SendDiscussMsg(discuss_id int64, message string) int32 {
	return int32(C.CQ_sendDiscussMsg(AUTOCODE, C.int64_t(discuss_id), gbk(message)))
}

func DeleteMsg(message_id int64) int32 {
	return int32(C.CQ_deleteMsg(AUTOCODE, C.int64_t(message_id)))
}

func SendLike(qq int64) int32 {
	return int32(C.CQ_sendLike(AUTOCODE, qq))
}

func SetGroupKick(group_id int64, qq int64, rejectaddrequest bool) int32 {
	return int32(C.CQ_setGroupKick(AUTOCODE, C.int64_t(group_id), C.int64_t(qq), C.int32_t(boolToInt32(rejectaddrequest))))
}

func SetGroupBan(group_id int64, qq int64, duration int64) int32 {
	return int32(C.CQ_setGroupBan(AUTOCODE, C.int64_t(group_id), C.int64_t(qq), C.int64_t(duration)))
}

func SetGroupAdmin(group_id int64, qq int64, isadmin bool) int32 {
	return int32(C.CQ_setGroupAdmin(AUTOCODE, C.int64_t(group_id), C.int64_t(qq), C.int32_t(boolToInt32(isadmin))))
}

func SetGroupWholeBan(group_id int64, enableban bool) int32 {
	return int32(C.CQ_setGroupWholeBan(AUTOCODE, C.int64_t(group_id), C.int32_t(boolToInt32(enableban))))
}

func SetGroupAnonymousBan(group_id int64, anomymous string, duration int64) int32 {
	return int32(C.CQ_setGroupAnonymousBan(AUTOCODE, C.int64_t(group_id), C.CString(anomymous), C.int64_t(duration)))
}

func SetGroupAnonymous(group_id int64, enableanomymous bool) int32 {
	return int32(C.CQ_setGroupAnonymous(AUTOCODE, C.int64_t(group_id), C.int32_t(boolToInt32(enableanomymous))))
}

func SetGroupCard(group_id int64, qq int64, newcard string) int32 {
	return int32(C.CQ_setGroupCard(AUTOCODE, C.int64_t(group_id), C.int64_t(qq), C.CString(newcard)))
}

func SetGroupLeave(group_id int64, isdismiss bool) int32 {
	return int32(C.CQ_setGroupLeave(AUTOCODE, C.int64_t(group_id), C.int32_t(boolToInt32(isdismiss))))
}

func SetGroupSpecialTitle(group_id int64, qq int64, title string, duration int64) int32 {
	return int32(C.CQ_setGroupSpecialTitle(AUTOCODE, C.int64_t(group_id), C.int64_t(qq), C.CString(title), C.int64_t(duration)))
}

func SetDiscussLeave(discuss_id int64) int32 {
	return int32(C.CQ_setDiscussLeave(AUTOCODE, C.int64_t(discuss_id)))
}

func SetFriendAddRequest(responseflag string, responseoperation int32, remark string) int32 {
	return int32(C.CQ_setFriendAddRequest(AUTOCODE, C.CString(responseflag), C.int32_t(responseoperation), C.CString(remark)))
}

func SetGroupAddRequest(responseflag string, requesttype int32, responseoperation int32, reason string) int32 {
	return int32(C.CQ_setGroupAddRequestV2(AUTOCODE, C.CString(responseflag),  C.int32_t(requesttype),  C.int32_t(responseoperation), C.CString(reason)))
}

func GetGroupMemberInfo(group_id int64, qq int64 , oncache bool) string {
	return C.GoString(C.CQ_getGroupMemberInfoV2(AUTOCODE, C.int64_t(group_id), C.int64_t(qq), C.int32_t(boolToInt32(oncache))))
}

func GetStrangerInfo(qq int64 , oncache bool) string {
	return C.GoString(C.CQ_getStrangerInfo(AUTOCODE, C.int64_t(qq), C.int32_t(boolToInt32(oncache))))
}

func GetCookies() string {
	return C.GoString(C.CQ_getCookies(AUTOCODE))
}

func GetCsrfToken() int32 {
	return int32(C.CQ_getCsrfToken(AUTOCODE))
}

func GetLoginQQ() int64 {
	return int64(C.CQ_getLoginQQ(AUTOCODE))
}

func GetLoginNick() string {
	return C.GoString(C.GoString(C.CQ_getLoginNick(AUTOCODE)))
}

func GetAppDirectory() string {
	return C.GoString(C.GoString(C.CQ_getAppDirectory(AUTOCODE)))
}

func SetFatal(message string) int32 {
	return int32(C.CQ_setFatal(AUTOCODE, C.CString(message)))
}

func CQ_getRecord(file string, outformat string) string {
	return C.GoString(AUTOCODE, C.CString(file), C.CString(outformat))
}

func gbk(c string) *C.char {
	return C.CString(mahonia.NewEncoder("gbk").ConvertString(c))
}

func boolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}