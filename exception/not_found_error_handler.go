package exception

type NotFoundErr struct {
	Message string
}

func (e NotFoundErr) Error() string {
	return e.Message
}

func NewNotFoundErr(message string) NotFoundErr {
	return NotFoundErr{Message: message}
}
