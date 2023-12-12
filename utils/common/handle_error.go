package common

type InvalidError struct {
	Message string
}

func (e InvalidError) Error() string {
	return e.Message
}
