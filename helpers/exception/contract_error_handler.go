package exception

// bad req
type BadRequestError struct {
	Message string
}

func (err BadRequestError) Error() string {
	return err.Message
}

// not found
type NotFoundError struct {
	Message string
}

func (err NotFoundError) Error() string {
	return err.Message
}
