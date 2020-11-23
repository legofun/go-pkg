package http

type HttpError struct {
	err
}

func (e *HttpError) Error() string {
	return
}
