package request

type CreateOrderReq struct {
	CustomerID    string          `json:"customer_id" swaggerignore:"true" `
	ServiceID     string          `json:"service_id" binding:"required" example:"OTO4"`
	Note          string          `json:"note" example:"Bác tài chờ 5 phút nhé"`
	PromotionCode []string        `json:"promotion_code" example:"FTR02"`
	OrderStage    []OrderStageReq `json:"order_stage"`
}

type OrderStageReq struct {
	Index    int    `json:"index" example:"1"`
	Location string `json:"location"  example:"9.688705,105.564740"`
	Address  string `json:"address"  example:"số 1 Võ Văn Ngân"`
}
