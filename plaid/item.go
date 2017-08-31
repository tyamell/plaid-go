package plaid

// CreateItem (POST /item/create) creates a new item and returns the item and access token.
//
// https://plaid.com/docs/api/#item-management
func (c *Client) CreateItem(username string, password string,
	institutionID string, products []string) (res *CreateItemResponse, err error) {
	return c.CreateItemWithOptions(username, password, institutionID, products, "")
}

// CreateItemWithOptions (POST /item/create) creates a new item and returns the item and access token
// with the optional Webhook parameter
//
// https://plaid.com/docs/api/#item-management
func (c *Client) CreateItemWithOptions(username string, password string,
	institutionID string, products []string, webhook string) (res *CreateItemResponse, err error) {
	var response CreateItemResponse
	request := createItemRequest{
		Credentials:   credentials{Username: username, Password: password},
		InstitutionID: institutionID,
		Products:      products,
		Options:       createItemOptions{Webhook: webhook},
	}
	err = c.post("/item/create", request, &response)
	return &response, err
}

// GetItem (POST /item/get) returns the Item associated with the provided access token.
//
// https://plaid.com/docs/api/#item-management
func (c *Client) GetItem(accessToken string) (res *ItemResponse, err error) {
	var response ItemResponse
	err = c.post("/item/get", itemRequest{AccessToken: accessToken}, &response)
	return &response, err
}

// DeleteItem (POST /item/delete) deletes the Item associated with the provided access token.
//
// https://plaid.com/docs/api/#item-management
func (c *Client) DeleteItem(accessToken string) (res *DeleteItemResponse, err error) {
	var response DeleteItemResponse
	err = c.post("/item/delete", itemRequest{AccessToken: accessToken}, &response)
	return &response, err
}

// UpdateWebhook (POST /item/webhook/update) updates the webhook for the Item associated with
// the provided access token.
//
// https://plaid.com/docs/api/#item-management
func (c *Client) UpdateWebhook(accessToken string, webhook string) (res *ItemResponse, err error) {
	var response ItemResponse
	req := updateWebhookRequest{
		AccessToken: accessToken,
		Webhook:     webhook,
	}
	err = c.post("/item/webhook/update", req, &response)
	return &response, err
}

// UpdateAccessTokenVersion (POST /item/access_token/update_version) updates a legacy access token to a V2 access token.
//
// https://plaid.com/docs/api/#item-management
func (c *Client) UpdateAccessTokenVersion(accessTokenV1 string) (res *UpdateTokenVersionResponse, err error) {
	var response UpdateTokenVersionResponse
	err = c.post("/item/access_token/update_version", updateTokenVersionRequest{AccessToken: accessTokenV1}, &response)
	return &response, err
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DeleteItemResponse is represents the response from DeleteItem
type DeleteItemResponse struct {
	Deleted   bool   `json:"deleted"`
	RequestID string `json:"request_id"`
}
type createItemOptions struct {
	Webhook string `json:"webhook,omitempty"`
}

type createItemRequest struct {
	Credentials   credentials       `json:"credentials"`
	InstitutionID string            `json:"institution_id"`
	Products      []string          `json:"initial_products"`
	Options       createItemOptions `json:"options,omitempty"`
}

type itemRequest struct {
	AccessToken string `json:"access_token"`
}

type updateTokenVersionRequest struct {
	AccessToken string `json:"access_token_v1"`
}

// UpdateTokenVersionResponse is represents the response from UpdateAccessTokenVersion
type UpdateTokenVersionResponse struct {
	AccessToken string `json:"access_token"`
	RequestID   string `json:"request_id"`
	ItemID      string `json:"item_id"`
}

type updateWebhookRequest struct {
	AccessToken string `json:"access_token"`
	Webhook     string `json:"webhook"`
}

// ItemResponse is represents the response the generic response from Item queries.
// It represents an item and request id
type ItemResponse struct {
	Item      Item
	RequestID string `json:"request_id"`
}

// CreateItemResponse is represents the response from CreateItem
type CreateItemResponse struct {
	AccessToken string `json:"access_token"`
	Item        Item
	MfaType     string
	Device      mfaDevice
	List        []mfaList
	Questions   []mfaQuestion
	Selections  []mfaSelection
	RequestID   string `json:"request_id"`
}

// Item is represents an item in the Plaid system
//
// https://plaid.com/docs/api/#item-management
type Item struct {
	AvailableProducts []string
	BillProducts      []string
	InstitutionID     string
	ItemID            string
	Webhook           string `json:"webhook"`
}
