package schema

import "sync"

// Declaring Models

type Album struct {
	ID     string  `json:"id" binding:"required"` // Declaring object property and ensuring that ID is required at all times
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type HTTPResponse struct {
	StatusCode int         `json:"status_code" binding:"required"`
	Message    string      `json:"message" binding:"required"`
	Data       interface{} `json:"data"`
}

// In-Memory store to be used to prevent multiple requests from disrupting
type DataStore struct {
	Albums []Album
	Mu     sync.Mutex
}
