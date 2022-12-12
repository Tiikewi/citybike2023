package types

type JSONResponse struct {
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

type JourneyRequest struct {
	DepTime      string `json:"departureTime" validate:"required"`
	RetTime      string `json:"returnTime" validate:"required"`
	DepStationId int    `json:"departureStationId" validate:"required"`
	RetStationId int    `json:"returnStationId" validate:"required"`
	Distance     int    `json:"distance" validate:"required,min=10"`
	Duration     int    `json:"duration" validate:"required,min=10"`
}

type Coordinates struct {
	X float64 `json:"x" validate:"required"`
	Y float64 `json:"y" validate:"required"`
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

type StationRequest struct {
	ID          int         `json:"id" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	Address     string      `json:"address" validate:"required"`
	City        string      `json:"city" validate:"required"`
	Operator    string      `json:"operator"`
	Capacity    int         `json:"capacity"`
	Coordinates Coordinates `json:"coordinates" validate:"required"`
}
