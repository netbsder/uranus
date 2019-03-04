package e

const (
	// global
	ERR_UNKNOWN = 99999
)

var MsgFlags = map[int]string{
	// global
	ERR_UNKNOWN: "unknown error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERR_UNKNOWN]
}
