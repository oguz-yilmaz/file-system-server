package fsmod

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
		StatusCode: 404, // todo: use a constant
		Message:    "File not found",
	}
}
