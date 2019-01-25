package err

var Msg = map[int]string{
	Success: "ok",
	Error:   "内部异常",
}

func GetMsg(code int) string {
	if msg, ok := Msg[code]; ok {
		return msg
	}

	return Msg[Error]
}
