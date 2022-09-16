package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Errno) Error() string {
	return err.Message
}

func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message + typed.Err.Error()
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return ErrInternalServer.Code, err.Error()
}
