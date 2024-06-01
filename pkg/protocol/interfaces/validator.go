package interfaces

type Validator interface {
	Validate(params ...any) error
}
