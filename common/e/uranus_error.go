package e

type UranusErr struct {
	Code int
	Msg  string
}

func NewError(code int) *UranusErr {
	return &UranusErr{Code: code, Msg: GetMsg(code)}
}

func NewErrorAll(code int, msg string) *UranusErr {
	return &UranusErr{Code: code, Msg: msg}
}

func (err *UranusErr) Error() string {
	return err.Msg
}
