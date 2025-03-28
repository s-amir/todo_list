package logs

type SimpleError struct {
	Output    string
	Operation string
}

func (s SimpleError) Error() string {
	return s.Output
}
