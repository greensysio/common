package res

// Response struct for return
type Response struct {
	Status  int                    `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}
