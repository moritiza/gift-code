package dto

// Report store received or fetched single report data
type Report struct {
	Mobile     string `json:"mobile"`
	Code       string `json:"code"`
	CodeCredit uint64 `json:"code_credit"`
}
