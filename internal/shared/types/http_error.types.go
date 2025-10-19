package types

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Message
}
