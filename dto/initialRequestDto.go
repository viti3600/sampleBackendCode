package dto

type GiftCardOrderRequest struct {
	SellerId      string `json:"sellerId"`
	BuyerId       string `json:"buyerId"`
	SkuId         int64  `json:"skuId"`
	Amount        int64  `json:"amount"`
	CorrelationId string `json:"correlationId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Recipient     string `json:"recipient"`
	Message       string `json:"message"`
}
