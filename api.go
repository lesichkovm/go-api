package github.com/lesichkovm/go-api

// Response defines an response for the API
type Response struct {
	Status  string                 `json:Status`
	Message string                 `json:Message`
	Data    map[string]interface{} `json:Data`
}

// Error returns an error message
func Error(message string) Response {
	return Response{Status: "error", Message: message, Data: map[string]interface{}{}}
}

// ErrorData returns an error message with data
func ErrorData(message string, data map[string]interface{}) Response {
	return Response{Status: "error", Message: message, Data: data}
}

// Success returns a success message
func Success(message string) Response {
	return Response{Status: "success", Message: message, Data: map[string]interface{}{}}
}

// SuccessData returns a success message with data
func SuccessData(message string, data map[string]interface{}) Response {
	return Response{Status: "success", Message: message, Data: data}
}
