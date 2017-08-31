package plaid_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tyamell/plaid-go/plaid"
	"github.com/tyamell/plaid-go/plaid/test"
)

type createLinkItemRequest struct {
	Credentials   credentials `json:"credentials"`
	InstitutionID string      `json:"institution_id"`
	Products      []string    `json:"initial_products"`
	PublicToken   string      `json:"public_key"`
}

type createLinkItemResponse struct {
	PublicToken string `json:"public_token"`
	Accounts    []plaid.Account
	RequestID   string `json:"request_id"`
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var _ = Describe("exchange-token", func() {
	var client *plaid.Client
	var publicToken string

	BeforeEach(func() {
		client = plaid.NewClient("5788670d66710877408d0687", "56226f1ef50ad084d6ea3bba62faae", plaid.Sandbox)
		var response createLinkItemResponse
		request := createLinkItemRequest{
			Credentials:   credentials{Username: test.GoodUser, Password: test.GoodPassword},
			InstitutionID: test.InstitutionId,
			Products:      test.Products,
			PublicToken:   "98c28dde1fa3b92fc54edefce04c54",
		}

		jsonBody, err := json.Marshal(request)
		if err != nil {
			Fail(err.Error())
		}

		err = client.PublicPost("/link/item/create", jsonBody, &response)
		if err != nil {
			Fail(err.Error())
		}
		publicToken = response.PublicToken
	})

	Describe("ExchangeToken", func() {
		It("returns public_token and access_token", func() {
			res, err := client.ExchangeToken(publicToken)
			Expect(err).To(BeNil(), "err should be nil")
			Expect(res.AccessToken).ToNot(BeEmpty(), "access token should not be empty")
		})
	})
})
