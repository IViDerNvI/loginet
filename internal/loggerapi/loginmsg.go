package loggerapi

import (
	"encoding/json"
	"regexp"

	myerr "github.com/IViDerNvI/loginet/internal/errors"
)

type logMsg struct {
	Result  int    `json:"result"`
	Msg     string `json:"msg"`
	RetCode int    `json:"ret_code"`
}

func newLogMsg(str string) *logMsg {
	var msg = &logMsg{
		Result:  -1,
		Msg:     "",
		RetCode: -1,
	}

	re := regexp.MustCompile(`\{.*\}`)
	matches := re.FindString(str)

	if len(matches) != 0 {
		json.Unmarshal([]byte(matches), msg)
	}

	return msg
}

func (lm *logMsg) getLoginStatus() error {
	switch lm.Msg {
	case "Portal协议认证成功！":
		return myerr.ErrSuccess
	case "AC999":
		return myerr.ErrDuplicate
	case "从Radius获取错误代码出现异常":
		return myerr.ErrIpError
	default:
		return myerr.ErrUnknown
	}
}

// Get Login Result
func GetLoginResult(response string) error {
	return newLogMsg(response).getLoginStatus()
}
