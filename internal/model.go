package internal

const (
	SuccessCode = 0
	SuccessMsg  = "success"
	ErrorCode   = 1
)

type Result struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}
