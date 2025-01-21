package common

const (
	IntervalError       = 500
	StatusNotFoundError = 404
	UnauthorizedError   = 401
	BadRequestError     = 400
	ForbiddenError      = 403
	ConnectionDBError   = 100001
)

var errCode = map[int]string{
	IntervalError:       "内部错误",
	StatusNotFoundError: "请求资源不存在",
	UnauthorizedError:   "未授权认证",
	BadRequestError:     "请求参数错误或格式不正确",
	ForbiddenError:      "请求拒绝, 权限不足",
	ConnectionDBError:   "数据源连通失败",
}

func Message(code int) string {
	if message, ok := errCode[code]; ok {
		return message
	}
	return errCode[IntervalError]
}
