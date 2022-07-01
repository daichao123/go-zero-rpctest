package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "success"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token 失效,请重新登录"
	message[TOKEN_GENERATE_ERROR] = "token 生成失败"
	message[DB_ERROR] = "数据库繁忙"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "修改行数未0"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
