package common

import "eastsunsoft.com/uranus-service/common/e"

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body"`
}

func NewRes() *Res {
	return NewResByParm(0, "", "")
}

func NewResByParm(code int, msg string, body interface{}) *Res {
	return &Res{code, msg, body}
}

func NewResErr(code int, msg string) *Res {
	return &Res{code, msg, ""}
}

func NewResByUranusErr(err *e.UranusErr) *Res {
	return &Res{err.Code, err.Msg, ""}
}

func (r *Res) WriteBody(b interface{}) *Res {
	r.Body = b
	return r
}
