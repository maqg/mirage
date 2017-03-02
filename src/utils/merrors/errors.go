package merrors

const (
	ERR_OCT_SUCCESS = iota
	ERR_DB_ERR
	ERR_NOT_ENOUGH_PARAS
	ERR_TOO_MANY_PARAS
	ERR_UNACCP_PARAS
	ERR_CMD_ERR
	ERR_COMMON_ERR
	ERR_SEGMENT_NOT_EXIST
	ERR_SEGMENT_ALREADY_EXIST
	ERR_TIMEOUT
	ERR_SYSCALL_ERR
	ERR_SYSTEM_ERR
	ERR_NO_SUCH_API
)

var GErrors = map[int]string{
	ERR_OCT_SUCCESS:           "Command Success",
	ERR_DB_ERR:                "Database Error",
	ERR_NOT_ENOUGH_PARAS:      "No Enough Paras",
	ERR_TOO_MANY_PARAS:        "Too Many Paras",
	ERR_UNACCP_PARAS:          "Unaccept Paras",
	ERR_CMD_ERR:               "Command Error",
	ERR_COMMON_ERR:            "Common Error",
	ERR_SEGMENT_NOT_EXIST:     "Segment Not Exist",
	ERR_SEGMENT_ALREADY_EXIST: "Segment Already Exist",
	ERR_TIMEOUT:               "Timeout Error",
	ERR_SYSCALL_ERR:           "System Call Error",
	ERR_SYSTEM_ERR:            "System Error",
	ERR_NO_SUCH_API:           "No Such API",
}

var GErrorsCN = map[int]string{
	ERR_OCT_SUCCESS:           "操作成功",
	ERR_DB_ERR:                "数据库错误",
	ERR_NOT_ENOUGH_PARAS:      "参数不足",
	ERR_TOO_MANY_PARAS:        "太多参数",
	ERR_UNACCP_PARAS:          "参数不合法",
	ERR_CMD_ERR:               "命令执行错误",
	ERR_COMMON_ERR:            "通用错误",
	ERR_SEGMENT_NOT_EXIST:     "对象不存在",
	ERR_SEGMENT_ALREADY_EXIST: "对象已存在",
	ERR_TIMEOUT:               "超时错误",
	ERR_SYSCALL_ERR:           "系统调用错误",
	ERR_SYSTEM_ERR:            "系统错误",
	ERR_NO_SUCH_API:           "无此API",
}

type MirageError struct {
	ErrorNo  int    `json:"no"`
	ErrorMsg string `json:"msg"`
}

func NewError(code int, message string) *MirageError {
	return &MirageError{
		ErrorNo:  code,
		ErrorMsg: message,
	}
}

func GetMsg(errorNo int) string {
	return GErrors[errorNo]
}

func GetMsgCN(errorNo int) string {
	return GErrorsCN[errorNo]
}
