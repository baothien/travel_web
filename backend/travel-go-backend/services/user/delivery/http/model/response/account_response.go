package response

type RequestRegisterRes struct {
	TransactionID string `json:"transaction_id"`
	ExpiredOTP    int    `json:"expired_otp"`
}
