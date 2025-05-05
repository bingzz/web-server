package schema

import "sync"

// Declaring Models

type Album struct {
	ID         int     `json:"id"` // Declaring object property and ensuring that ID is required at all times
	Title      string  `json:"title"`
	AuthorID   int     `json:"author_id"`
	AuthorName string  `json:"author_name"`
	Price      float32 `json:"price"`
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

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` // hashed password
	Name     string `json:"name"`
	Status   bool   `json:"status"`
}
