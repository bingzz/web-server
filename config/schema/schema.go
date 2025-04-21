package schema

type ServerConfig struct {
	port        string
	url         string
	environment string
}

type RequestPayload struct {
	Name  string `json:"name"`  // Exported (uppercase) with JSON tag
	Email string `json:"email"` // Exported (uppercase) with JSON tag
}
