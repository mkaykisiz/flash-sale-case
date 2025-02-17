package constants

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	INVALIDPARAMS: "invalid",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
