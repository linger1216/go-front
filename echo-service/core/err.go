package core

var (
// StatusInternalServerError = model.NewError(http.StatusInternalServerError, "Status Internal Server Error", "")
)

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}