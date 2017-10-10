package api

type ErrMessageInterface interface {
	GetException() string
	GetMessage() string
}

type ErrMessage struct {
	Message   string
	Exception string
}

func (e *ErrMessage) GetException() string {
	return e.Exception
}

func (e *ErrMessage) GetMessage() string {
	return e.Message
}
