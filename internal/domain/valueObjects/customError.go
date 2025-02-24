package valueobjects

type CustomError struct {
	Code string
	Desc string
}

func NewCustomError(code, desc string) error {
	return CustomError{code, desc}
}

func (ref CustomError) Error() string {
	return ref.Desc
}
