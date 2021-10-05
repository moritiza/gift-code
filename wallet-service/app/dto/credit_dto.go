package dto

// GetCredit store received data for getting user credit amount
type GetCredit struct {
	Mobile string `json:"mobile" validate:"required,mobile"`
	Credit uint64 `json:"credit,omitempty"`
}

// SetDiscountCredit store received data for setting user discount credit amount
type SetDiscountCredit struct {
	Mobile string `json:"mobile" validate:"required,mobile"`
	Credit uint64 `json:"credit" validate:"required,credit"`
}
