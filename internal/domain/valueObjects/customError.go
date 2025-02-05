package valueobjects

type CustomError struct {
	Code string
	Desc string
}

func New(code, desc string) error {
	return CustomError{code, desc}
}

func (ref CustomError) Error() string {
	return ref.Desc
}
