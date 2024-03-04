package dto

type ProductsResponse struct {
	ResponseCode    string           `json:"responseCode"`
	ResponseMessage string           `json:"responseMessage"`
	PayLoad         []PayLoadDetails `json:"payLoad"`
}
type PayLoadDetails struct {
	SKUID                    int                 `json:"skuId"`
	ProductID                int                 `json:"productId"`
	ProductName              string              `json:"productName"`
	Category                 string              `json:"category"`
	IsSalesTaxCharged        bool                `json:"isSalesTaxCharged"`
	CountryCode              string              `json:"countryCode"`
	BenefitType              string              `json:"benefitType"`
	Validity                 string              `json:"validity"`
	ProductDesc              string              `json:"productDescription"`
	LocalPhoneNumberLength   int                 `json:"localPhoneNumberLength"`
	InternationalCountryCode interface{}         `json:"internationalCountryCode"`
	AllowDecimal             bool                `json:"allowDecimal"`
	Fee                      float64             `json:"fee"`
	OperatorId               int                 `json:"operatorId"`
	OperatorName             string              `json:"operatorName"`
	ImageUrl                 string              `json:"imageUrl"`
	AdditionalInformation    string              `json:"additionalInformation"`
	MinFaceValue             MinFaceValueDetails `json:"min"`
	MaxFaceValue             MaxFaceValueDetails `json:"max"`
}
type MinFaceValueDetails struct {
	FaceValue            float64 `json:"faceValue"`
	FaceValueCurrency    string  `json:"faceValueCurrency"`
	DeliveredAmount      float64 `json:"deliveredAmount"`
	DeliveryCurrencyCode string  `json:"deliveryCurrencyCode"`
	Cost                 float64 `json:":cost"`
	CostCurrency         string  `json:"costCurrency"`
}

type MaxFaceValueDetails struct {
	FaceValue            float64 `json:"faceValue"`
	FaceValueCurrency    string  `json:"faceValueCurrency"`
	DeliveredAmount      float64 `json:"deliveredAmount"`
	DeliveryCurrencyCode string  `json:"deliveryCurrencyCode"`
	Cost                 float64 `json:":cost"`
	CostCurrency         string  `json:"costCurrency"`
}
