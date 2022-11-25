package types

type PingResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
