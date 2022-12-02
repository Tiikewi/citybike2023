package types

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

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Station struct {
	FID         int         `json:"fid"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	Operator    string      `json:"operator"`
	Capacity    int         `json:"capacity"`
	Coordinates Coordinates `json:"coordinates"`
	Returns     int         `json:"returns"`
	Departures  int         `json:"departures"`
}
