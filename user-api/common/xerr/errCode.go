package xerr

//OK 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//SERVER_COMMON_ERROR 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001           //系统公告错误
const REUQEST_PARAM_ERROR uint32 = 100002           //请求参数错误
const TOKEN_EXPIRE_ERROR uint32 = 100003            //token 超时错误
const TOKEN_GENERATE_ERROR uint32 = 100004          //token 生成错误
const DB_ERROR uint32 = 100005                      //数据库错误
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006 //数据库修改错误
