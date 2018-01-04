package plaid

import "time"

// GetTransactions returns the accounts and transactions associated with the provided
// access token. Transactions are provided within the range of Start Date and End Date, by default
// count is 250 and offset is 0.
func (c *Client) GetTransactions(accessToken string, startDate time.Time,
	endDate time.Time, count int, offset int) (res *GetTransactionsResponse, err error) {
	if count == 0 {
		count = 250
	}
	request := getTransactionsRequest{
		AccessToken: accessToken,
		Options:     getTransactionsOptions{count, offset},
		StartDate:   startDate.Format("2006-01-02"),
		EndDate:     endDate.Format("2006-01-02"),
	}

	var response GetTransactionsResponse
	err = c.post("/transactions/get", request, &response)
	return &response, err
}

type getTransactionsOptions struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
}

type getTransactionsRequest struct {
	AccessToken string                 `json:"access_token"`
	StartDate   string                 `json:"start_date"`
	EndDate     string                 `json:"end_date"`
	Options     getTransactionsOptions `json:"options"`
}

// GetTransactionsResponse represents the response from GetTransactions
type GetTransactionsResponse struct {
	Accounts          []Account     `json:"accounts"`
	TotalTransactions int           `json:"total_transactions"`
	Transactions      []Transaction `json:"transactions"`
	Item              Item          `json:"item"`
	RequestID         string        `json:"request_id"`
}

// Transaction is returned by GetTransactions and represents data for credit and depository-type Accounts.
// Transaction data is standardized across financial institutions,
// and in many cases transactions are linked to a clean name, entity type,
// location, and category. Similarly, account data is standardized and returned
// with a clean name, number, balance, and other meta information where available.
// TODO make TransactionType an enum
//
// https://plaid.com/docs/api/#transactions
type Transaction struct {
	TransactionID        string   `json:"transaction_id"`
	AccountID            string   `json:"account_id"`
	Categories           []string `json:"category"`
	CategoryID           string   `json:"category_id"`
	TransactionType      string   `json:"transaction_type"`
	Amount               float64
	Date                 *PlaidTime
	Location             *Location
	PaymentMetadata      *PaymentMetadata `json:"payment_meta"`
	Pending              bool
	PendingTransactionID string `json:"pending_transaction_id"`
	AccountOwner         string `json:"account_owner"`
}

// Location is detailed merchant location data including address, city,
// state, zip, lat (latitude), and lon (longitude) where available.
type Location struct {
	Address   string
	City      string
	State     string
	Zip       string
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// PaymentMetadata is detailed payment and payment processor data
// including reference_number, ppd_id, and payee_name where available.
type PaymentMetadata struct {
	ReferenceNumber        string `json:"reference_number"`
	PaymentProcessorDataID string `json:"ppd_id"`
	PayeeName              string `json:"payee_name"`
	Payer                  string `json:"payer"`
	PaymentMethod          string `json:"payment_method"`
	PaymentProcessor       string `json:"payment_processor"`
	Reason                 string `json:"reason"`
	ByOrderOf              string `json:"by_order_of"`
}
