package exception

type NotFoundErr struct {
	Error string
}

func NotFoundErrHandler(error string) NotFoundErr {
	return NotFoundErr{Error: error}
}
