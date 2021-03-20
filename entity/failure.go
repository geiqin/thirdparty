package entity

//失败信息
type Failure struct {
	code int
	msg  string
	err  error
}

func NewFailure(code int, msg string) *Failure {
	return &Failure{
		code: code,
		msg:  msg,
	}
}

func (t *Failure) copy() *Failure {
	return NewFailure(t.code, t.msg)
}

func (t *Failure) WithErr(err error) *Failure {
	r := t.copy()
	r.err = err
	r.WithMsg(err.Error())
	return r
}

func (t *Failure) WithMsg(msg string) *Failure {
	r := t.copy()
	if r.msg == "" {
		r.msg = msg
	} else {
		r.msg += ":" + msg
	}
	return r
}
