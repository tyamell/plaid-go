// Package plaid implements a Go client for the Plaid API (https://plaid.com/docs)
package plaid

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NewClient instantiates a Client associated with a client id, secret and environment.
// See https://plaid.com/docs/api/#gaining-access.
func NewClient(clientID string, secret string, environment EnvironmentURL) *Client {
	return &Client{clientID, secret, environment, &http.Client{}}
}

// NewCustomClient is the as above but with additional parameter to pass http.Client. This is required
// if you want to run the code on Google AppEngine which prohibits use of http.DefaultClient
func NewCustomClient(clientID string, secret string, environment EnvironmentURL,
	httpClient *http.Client) *Client {
	return &Client{clientID, secret, environment, httpClient}
}

// Client is only exported for method documentation purposes.
// Instances should only be created through the 'NewClient' function.
//
// See https://github.com/golang/go/issues/7823.
type Client struct {
	clientID    string
	secret      string
	environment EnvironmentURL
	httpClient  *http.Client
}

type EnvironmentURL string

// Sandbox is the URL for the Sandbox environment at Plaid
// Use this for integration tests
var Sandbox EnvironmentURL = "https://sandbox.plaid.com"

// Development is the URL for the Development environment at Plaid
// Use this for while developing your application
var Development EnvironmentURL = "https://development.plaid.com"

// Production is the URL for the Production environment at Plaid
// Use this when you application is deployed to production
var Production EnvironmentURL = "https://api.plaid.com"

// Account resprests the account data in the Plaid system
// TODO make AccountType and SubAccountType enums
//
// https://plaid.com/docs/api/#accounts
type Account struct {
	ID             string `json:"account_id"`
	AccountType    string `json:"type"`
	OfficialName   string `json:"official_name"`
	Name           string
	Mask           string
	ItemID         string `json:"item_id"`
	InstitutionID  string `json:"institution_id"`
	SubAccountType string `json:"subtype"`
	Balance        struct {
		Available float64 `json:"available"`
		Current   float64 `json:"current"`
		Limit     float64 `json:"limit"`
	} `json:"balances"`
}

func (c *Client) post(endpoint string, body interface{}, responseStructure interface{}) error {
	bodyBytes, err := c.addSecretToBody(body)
	if err != nil {
		return nil
	}
	return c.httpCall(http.MethodPost, endpoint, bodyBytes, &responseStructure)
}

func (c *Client) addSecretToBody(body interface{}) ([]byte, error) {
	authBody := map[string]interface{}{}
	bodyString, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyString, &authBody)
	if err != nil {
		return nil, err
	}
	authBody["client_id"] = c.clientID
	authBody["secret"] = c.secret
	jsonBody, err := json.Marshal(authBody)
	// fmt.Println(string(jsonBody))
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func (c *Client) PublicPost(endpoint string, body []byte, responseStructure interface{}) error {
	return c.httpCall(http.MethodPost, endpoint, body, &responseStructure)
}

func (c *Client) httpCall(method string, endpoint string, body []byte, responseStructure interface{}) error {
	request, err := http.NewRequest(method, string(c.environment)+endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "plaid-go")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println(string(raw))
	fmt.Println()

	// fmt.Printf("Printing results for %s\n", string(c.environment)+endpoint)
	// fmt.Println(string(raw))

	err = response.Body.Close()
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK:
		if err = json.Unmarshal(raw, responseStructure); err != nil {
			return err
		}
		return nil
	default:
		if response.StatusCode >= 400 {
			var plaidErr plaidError
			fmt.Println(string(raw))
			if err = json.Unmarshal(raw, &plaidErr); err != nil {
				return err
			}
			plaidErr.StatusCode = response.StatusCode
			return plaidErr
		}
		return errors.New("unknown Plaid Error - Status:" + string(response.StatusCode))
	}
}
