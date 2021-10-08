package model

const (
	ErrorTypeFatal        = "Fatal"
	ErrorTypeError        = "Error"
	ErrorTypeValidation   = "Validation"
	ErrorTypeInfo         = "Info"
	ErrorTypeDebug        = "Debug"
	ErrorTypeUnauthorized = "Unauthorized"
)

type ErrorDetail struct {
	ErrType string `json:"err_type,omitempty"`
	ErrMsg  string `json:"err_msg,omitempty"`
}

type Response struct {
	Data    interface{}   `json:"data,omitempty"`
	Status  int           `json:"status,omitempty"`
	Error   []ErrorDetail `json:"error,omitempty"`
	Message string        `json:"message,omitempty"`
}
