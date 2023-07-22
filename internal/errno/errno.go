package errno

type ErrorNo struct {
	Err error
	No  int
}

func (e ErrorNo) Error() string {
	if e.No > 0 {
		return ErrMsgMap[e.No]
	}
	return e.Err.Error()
}

func NewErrorNo(err error, no int) ErrorNo {
	return ErrorNo{
		Err: err,
		No:  no,
	}
}
