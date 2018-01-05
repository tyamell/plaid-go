package plaid

// import (
// 	"bytes"
// 	"encoding/json"
// )
// GetTransactions returns the accounts and transactions associated with the provided
// access token. Transactions are provided within the range of Start Date and End Date, by default
// count is 250 and offset is 0.
func (c *Client) GetBalance(accessToken string) (res *GetBalanceResponse, err error) {
	request := getBalanceRequest{
		AccessToken: accessToken,
		AccountIds:  nil,
	}

	var response GetBalanceResponse
	err = c.post("/accounts/balance/get", request, &response)
	return &response, err
}

type getBalanceRequest struct {
	AccessToken string `json:"access_token"`
	Options     struct {
		AccountIds []string `json:"account_ids,omitempty"`
	} `json:"options, omitempty"`
}

// GetTransactionsResponse represents the response from GetTransactions
type GetBalanceResponse struct {
	Accounts  []Account `json:"accounts"`
	Item      Item      `json:"item"`
	RequestID string    `json:"request_id"`
}

// // Balance (POST /balance) retrieves real-time balance for a given access token.
// //
// // See https://plaid.com/docs/api/#balance.
// func (c *Client) Balance(accessToken string) (postRes *postResponse, err error) {
// 	jsonText, err := json.Marshal(balanceJson{
// 		ClientID:    c.clientID,
// 		Secret:      c.secret,
// 		AccessToken: accessToken,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	postRes, _, err = c.postAndUnmarshal("/balance", bytes.NewReader(jsonText))
// 	return postRes, err
// }

// type balanceJson struct {
// 	ClientID    string `json:"client_id"`
// 	Secret      string `json:"secret"`
// 	AccessToken string `json:"access_token"`
// }
