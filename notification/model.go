package notification

// Payload is mandatory information for notification
type Payload struct {
	UserID      string                 `json:"user_id,omitempty" valid:"required"`
	UserRole    string                 `json:"user_role,omitempty" valid:"required"`
	SenderId    string                 `json:"sender_id"`
	SenderRole  string                 `json:"sender_role"`
	Locale      string                 `json:"locale,omitempty"`
	Email       string                 `json:"email"`
	Message     string                 `json:"message,omitempty" valid:"required"`
	Feature     string                 `json:"feature,omitempty" valid:"required"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
	TimeStr     string                 `json:"time"`
}
