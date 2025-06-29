package model

type Request struct {
	Complaint string `json:"complaint" binding:"required"` // binding for Gin validation
}

type DiseasePrediction struct {
	Disease    string  `json:"disease"`    // predicted disease name
	Confidence float64 `json:"confidence"` // confidence score (0-1)
}

type Response struct {
	Predictions []DiseasePrediction `json:"predictions"`
}
