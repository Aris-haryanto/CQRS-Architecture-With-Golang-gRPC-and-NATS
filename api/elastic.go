package api

type ElDeposit struct {
	Amount  int64  `json:"amount"`
	From    string `json:"from"`
	Approve int8   `json:"approve"`
}
