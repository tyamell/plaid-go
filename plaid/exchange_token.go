package plaid

// ExchangeToken (POST /item/public_token/exchange) exchanges a public token for an access token.
//
// https://plaid.com/docs/api/#exchange-token-flow
func (c *Client) ExchangeToken(publicToken string) (res *ExchangeTokenResponse, err error) {
	var response ExchangeTokenResponse
	var request = exchangeTokenRequest{
		PublicToken: publicToken,
	}

	err = c.post("/item/public_token/exchange", request, &response)
	return &response, err
}

type exchangeTokenRequest struct {
	PublicToken string `json:"public_token"`
}

// ExchangeTokenResponse is a struct representing the response from ExchangeToken
type ExchangeTokenResponse struct {
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
	RequestID   string `json:"request_id"`
}
