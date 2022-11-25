package types

type PingResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Journey struct {
	ID             int    `json:"id"`
	DepTime        string `json:"departureTime"`
	RetTime        string `json:"returnTime"`
	DepStationId   int    `json:"departureStationId"`
	DepStationName string `json:"departureStationName"`
	RetStationId   int    `json:"returnStationId"`
	RetStationName string `json:"returnStationName"`
	Distance       int    `json:"distance"`
	Duration       int    `json:"duration"`
}
