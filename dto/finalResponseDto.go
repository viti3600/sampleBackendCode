package dto

type GetAllGiftCardResponse struct {
	PPNGCId           int64   `json:"ppnGCId"`
	SellerId          string  `json:"sellerId"`
	BuyerId           string  `json:"buyerId"`
	FirstName         string  `json:"firstName"`
	LastName          string  `json:"lastName"`
	ResponseCode      string  `json:"responseCode"`
	ResponseMessage   string  `json:"responseMessage"`
	CorrelationId     string  `json:"correlationId"`
	TransactionId     int64   `json:"transactionId"`
	TransactionDate   string  `json:"transactionDate"`
	InvoiceAmount     float64 `json:"invoiceAmount"`
	FaceValue         float64 `json:"faceValue"`
	Discount          float64 `json:"discount"`
	Fee               float64 `json:"fee"`
	SkuId             int64   `json:"skuId"`
	ProductName       string  `json:"productName"`
	Instructions      string  `json:"instructions"`
	TopupDetail       string  `json:"topupDetail"`
	TopupPins         string  `json:"topupPins"`
	CardNumber        string  `json:"cardNumber"`
	CardPIN           string  `json:"pin"`
	CertificateLink   string  `json:"certificateLink"`
	SimInfo           string  `json:"simInfo"`
	BillPaymentDetail string  `json:"billPaymentDetail"`
}
