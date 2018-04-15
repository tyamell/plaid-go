package plaid

import (
	"github.com/pkg/errors"
)

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// )

// GetInstitution returns information for a single institution given an ID.
func (c *Client) GetInstitution(id string) (*GetInstitutionResponse, error) {
	if id == "" {
		return nil, errors.New("institution id must be specified")
	}

	request := getInstitutionRequest{
		InstitutionID: id,
	}

	var response GetInstitutionResponse
	err := c.postPublic("/institutions/get_by_id", request, &response)
	return &response, errors.Wrap(err, "error calling /institutions/get_by_id")
}

// GetAllInstitutions returns all institutions based on the product array provided. A nil product
// will return all institutions
func (c *Client) GetAllInstitutions(count, offset int, products []string) (*GetAllInstitutionResponse, error) {
	if count > 500 {
		return nil, errors.New("count must be less than or equal to 500")
	}
	request := getAllInstitutionsRequest{
		Count:  count,
		Offset: offset,
	}
	if products != nil {
		request.Options.Products = products
	}
	var response GetAllInstitutionResponse
	err := c.postPublic("/institutions/get", request, &response)
	return &response, errors.Wrap(err, "error calling /institutions/get")
}

// SearchInstitutions returns details for all institutions that match the query parameters.
// Use the product parameter to filter the Institutions based on whether they support all products listed in products.
// Provide nil to get institutions regardless of supported products.
func (c *Client) SearchInstitutions(query string, products []string) (*SearchInstitutionResponse, error) {
	if query == "" {
		return nil, errors.New("a query must be provided")
	}

	request := searchInstitutionRequest{
		Query:    query,
		Products: products,
	}
	var response SearchInstitutionResponse
	err := c.postPublic("/institutions/search", request, &response)
	return &response, errors.Wrap(err, "error calling /institutions/search")
}

type searchInstitutionRequest struct {
	Query    string   `json:"query"`
	Products []string `json:"products"`
}

type getInstitutionRequest struct {
	InstitutionID string `json:"institution_id"`
}

type getAllInstitutionsRequest struct {
	Count   int `json:"count"`
	Offset  int `json:"offset"`
	Options struct {
		Products []string `json:"products"`
	} `json:"options"`
}

type SearchInstitutionResponse struct {
	Institutions []Institution `json:"institutions"`
	RequestID    string        `json:"request_id"`
}

type GetAllInstitutionResponse struct {
	Institutions []Institution `json:"institutions"`
	Total        int           `json:"total"`
	RequestID    string        `json:"request_id"`
}

type GetInstitutionResponse struct {
	Institution Institution `json:"institution"`
	RequestID   string      `json:"request_id"`
}

type InstitutionCredentials struct {
	Label string `json:"label"` // e.g.: "Password"
	Name  string `json:"name"`  // e.g.: "PIN"
	Type  string `json:"type"`  // e.g.: "Online ID"
}

type Institution struct {
	Credentials []InstitutionCredentials `json:"credentials"`
	Name        string                   `json:"name"`           // e.g.: "Bank of America"
	HasMFA      bool                     `json:"has_mfa"`        // e.g.: true
	ID          string                   `json:"institution_id"` // e.g.: "5301a93ac140de84910000e0"
	MFA         []string                 `json:"mfa"`            // e.g.: ["code", "list", "questions"]
	Products    []string                 `json:"products"`       // e.g.: ["connect", "auth", "balance"]
}
