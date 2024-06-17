package fsmod

const (
	GenericErrorCode         = 100
	StatusNotFoundCode       = 404
	StatusInvalidRequestCode = 400
)

var errorMessages = map[int]string{
	GenericErrorCode:         "something went wrong",
	StatusNotFoundCode:       "file not found",
	StatusInvalidRequestCode: "invalid parameters",
}

// FileNotFoundError is an error type for file not found
type FileNotFoundError struct {
	FileName   string
	StatusCode int
	Message    string
}

func (e *FileNotFoundError) Error() string {
	return e.Message
}

func NewFileNotFoundError(fileName string) *FileNotFoundError {
	return &FileNotFoundError{
		FileName:   fileName,
		StatusCode: StatusNotFoundCode,
		Message:    errorMessages[StatusNotFoundCode],
	}
}

// InvalidParamsError is an error type for invalid parameters
type InvalidParamsError struct {
	StatusCode int
	Message    string
}

func (e *InvalidParamsError) Error() string {
	return e.Message
}

func NewInvalidParamsError() *InvalidParamsError {
	return &InvalidParamsError{
		StatusCode: StatusInvalidRequestCode,
		Message:    errorMessages[StatusInvalidRequestCode],
	}
}

// GenericError is an error type for generic errors
type GenericError struct {
	StatusCode int
	Message    string
}

func (e *GenericError) Error() string {
	return e.Message
}

func NewGenericError() *GenericError {
	return &GenericError{
		StatusCode: GenericErrorCode,
		Message:    errorMessages[GenericErrorCode],
	}
}
