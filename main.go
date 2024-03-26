package main

type BalanceResponse struct {
	Total int    `json:"total"`
	Date  string `json:"balance"`
	Limit string `json:"limit"`
}

type TransactionResponse struct {
	Value       int
	Type        string
	Description string
	Date        string
}

type Response struct {
	Balance             BalanceResponse       `json:"balance"`
	TransactionResponse []TransactionResponse `json:"lastTransactions"`
}
type Balance struct {
	Limit   int    `json:"limit"`
	Balance string `json:"balance"`
}

type TransactionRequest struct {
	Value       int    `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
}


func main() {

}
