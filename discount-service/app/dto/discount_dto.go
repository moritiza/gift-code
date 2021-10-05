package dto

// Discount store received discount data
type Discount struct {
	Code       string `json:"code" validate:"required,code"`
	CodeCredit uint64 `json:"code_credit" validate:"required,code_credit"`
	Count      uint   `json:"count" validate:"required,count"`
}

// SubmitDiscount store enterded dicount code
type SubmitDiscount struct {
	Mobile string `json:"mobile" validate:"required,mobile"`
	Code   string `json:"code" validate:"required,code"`
}

// SetDiscountCredit store data for setting user discount credit amount into wallet service
type SetDiscountCredit struct {
	Mobile string `json:"mobile"`
	Credit uint64 `json:"credit"`
}
