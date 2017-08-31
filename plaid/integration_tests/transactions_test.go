package plaid_test

import (
	"time"

	"github.com/tyamell/plaid-go/plaid"
	"github.com/tyamell/plaid-go/plaid/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("transaction", func() {

	var client *plaid.Client

	BeforeEach(func() {
		client = plaid.NewClient("5788670d66710877408d0687", "56226f1ef50ad084d6ea3bba62faae", plaid.Sandbox)
	})

	Describe("GetTransactions", func() {
		It("returns all transactions", func() {
			res, err := client.GetTransactions(test.AccessToken, time.Now().AddDate(0, -1, 0), time.Now(), 100, 0)
			Expect(err).To(BeNil(), "err should be nil")
			Expect(res.Item).ToNot(BeNil(), "item should not be nil")
			Expect(res.Transactions).ToNot(BeEmpty(), "transactions should not be empty")
			Expect(res.Accounts).ToNot(BeEmpty(), "accounts token should not be empty")
		})

	})
})
