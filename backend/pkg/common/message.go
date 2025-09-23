package common

const (
	SUCCESS_OK     = 200
	INVALID_PARAMS = 400
	ERROR_GENERIC  = 500

	ERROR_USER_CREDENTIALS_INVALID = 1001
	ERROR_USER_MISSING_TOKEN       = 1002
	ERROR_USER_TOKEN_EXPIRED       = 1003
	ERROR_USER_TOKEN_INVALID       = 1004

	ERROR_INVALID_PARAMS = 1500

	ERROR_PAYMENT_NOT_FOUND   = 2001
	ERROR_PAYMENT_COUNT_FAIL  = 2002
	ERROR_PAYMENT_REVIEW_FAIL = 2003

	ERROR_INSUFFICIENT_PERMISSIONS = 3001
	ERROR_MISMATCHED_ROLE          = 3002
)

var MsgFlags = map[int]string{
	SUCCESS_OK:     "ok",
	INVALID_PARAMS: "invalid params",
	ERROR_GENERIC:  "unknown error",

	ERROR_INVALID_PARAMS: "invalid parameters",

	ERROR_PAYMENT_NOT_FOUND:   "payment record not found",
	ERROR_PAYMENT_COUNT_FAIL:  "failed to get total payment records",
	ERROR_PAYMENT_REVIEW_FAIL: "failed to review payment; wrong action or payment already reviewed",

	ERROR_USER_CREDENTIALS_INVALID: "wrong username or password",
	ERROR_USER_MISSING_TOKEN:       "missing jwt access token",

	ERROR_USER_TOKEN_EXPIRED: "jwt access token expired",
	ERROR_USER_TOKEN_INVALID: "invalid jwt access token",

	ERROR_INSUFFICIENT_PERMISSIONS: "insufficient permissions",
	ERROR_MISMATCHED_ROLE:          "mismatched user role",
}

func MessageCode(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR_GENERIC]
}
