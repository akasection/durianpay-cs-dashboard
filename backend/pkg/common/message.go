package common

const (
	SUCCESS_OK     = 200
	INVALID_PARAMS = 400
	ERROR_GENERIC  = 500

	ERROR_USER_CREDENTIALS_INVALID = 1001

	ERROR_PAYMENT_NOT_FOUND   = 2001
	ERROR_PAYMENT_COUNT_FAIL  = 2002
	ERROR_PAYMENT_REVIEW_FAIL = 2003
)

var MsgFlags = map[int]string{
	SUCCESS_OK:     "ok",
	INVALID_PARAMS: "invalid params",
	ERROR_GENERIC:  "unknown error",

	ERROR_USER_CREDENTIALS_INVALID: "wrong username or password",
}

func MessageCode(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR_GENERIC]
}
