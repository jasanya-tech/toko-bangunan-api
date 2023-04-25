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

// unauthorize
type Unauthorize struct {
	Message string
}

func (err Unauthorize) Error() string {
	return err.Message
}

// UnprocessableEntity
type UnprocessableEntity struct {
	Message string
}

func (err UnprocessableEntity) Error() string {
	return err.Message
}

// Forbidden
type Forbidden struct {
	Message string
}

func (err Forbidden) Error() string {
	return err.Message
}
