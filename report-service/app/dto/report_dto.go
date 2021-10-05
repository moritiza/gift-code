package dto

// Reports store all reports to return data
type Reports struct {
	Reports []Report `json:"reports"`
}

// Report store received or fetched single report data
type Report struct {
	Mobile     string `json:"mobile" validate:"required,mobile"`
	Code       string `json:"code" validate:"required,code"`
	CodeCredit uint64 `json:"code_credit" validate:"required,code_credit"`
}
