package error

// VKError permorms special type of error which contains code and description:
// same as native VK.com error format
type VKError struct {
	Code string `json:"error"`
	Description string `json:"error_description"`
}

// Check interface
var _ error = &VKError{}

// Error implements error interface for VKError.
func (e *VKError) Error() string {
	return e.Description
}

// New creates a new VKError by the given code and description
func New(code string, description string) *VKError {
	return &VKError{Code: code, Description: description}
}
