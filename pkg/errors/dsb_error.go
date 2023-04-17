package errors

type DSBError struct {
	Message string
}

func NewDSB(message string) DSBError {
	return DSBError{
		Message: message,
	}
}

func (e DSBError) Error() string {
	return e.Message
}

func (e DSBError) DSBMessage() string {
	return e.Message
}
