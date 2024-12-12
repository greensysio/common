package sms

type (
	// VerifyCodeInfo : needed data for sending SMS
	VerifyCodeInfo struct {
		Phone       string
		CountryCode string
		VerifyCode  string
	}

	SMSRequest struct {
		Phone       string
		CountryCode string
		Content     string
	}

	// Payload : needed data
	Payload struct {
		To      []string `json:"to"`
		Content string   `json:"content"`
		SmsType string   `json:"sms_type"`
		Sender  string   `json:"sender"`
	}

	// DataRes : Data Response
	DataRes struct {
		Status  string `json:"status"`
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			TranID       int      `json:"tranId"`
			TotalSMS     int      `json:"totalSMS"`
			TotalPrice   int      `json:"totalPrice"`
			InvalidPhone []string `json:"invalidPhone"`
		}
	}
)
