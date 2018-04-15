package plaid

import "github.com/pkg/errors"

// GetCategories returns detailed information on categories returned by Plaid.
func (c *Client) GetCategories() (*GetCategoriesResponse, error) {
	var response GetCategoriesResponse
	err := c.PublicPost("/categories/get", []byte("{}"), &response)
	return &response, errors.Wrap(err, "error calling /categories/get")
}

type GetCategoriesResponse struct {
	Categories []Category `json:"hierarchy"`
	RequestID  string     `json:"request_id"`
}

type Category struct {
	Hierarchy  []string `json:"hierarchy"`
	Group      string   `json:"group"`
	CategoryID string   `json:"category_id"`
}
