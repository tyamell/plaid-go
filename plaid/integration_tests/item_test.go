package plaid_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tyamell/plaid-go/plaid"
	"github.com/tyamell/plaid-go/plaid/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("item", func() {
	var client *plaid.Client

	BeforeEach(func() {
		client = plaid.NewClient("5788670d66710877408d0687", "56226f1ef50ad084d6ea3bba62faae", plaid.Sandbox)
	})

	Describe("Creating items", func() {
		It("returns access_token", func() {
			res, err := client.CreateItem(test.GoodUser, test.GoodPassword, test.InstitutionId, test.Products)
			Expect(err).To(BeNil(), "err should be nil")
			Expect(res.Item).ToNot(BeNil(), "item should not be nil")
			Expect(res.MfaType).To(BeEmpty(), "mfa type should be blank")
			Expect(res.AccessToken).ToNot(BeEmpty(), "access token should not be empty")
		})

		It("returns item with webhook", func() {
			res, err := client.CreateItemWithOptions(test.GoodUser, test.GoodPassword,
				test.InstitutionId, test.Products, "http://test.foo.com")
			Expect(err).To(BeNil(), "err should be nil")
			Expect(res.Item).ToNot(BeNil(), "item should not be nil")
			Expect(res.Item.Webhook).To(Equal("http://test.foo.com"), "webhook should be http://test.foo.com")
			Expect(res.MfaType).To(BeEmpty(), "mfa type should be blank")
			Expect(res.AccessToken).ToNot(BeEmpty(), "access token should not be empty")
		})
	})

	Describe("GetItem", func() {
		It("returns item with webhook", func() {
			res, err := client.GetItem(test.AccessToken)
			Expect(err).To(BeNil(), "err should be nil")
			Expect(res.Item).ToNot(BeNil(), "item should not be nil")
			Expect(res.Item.Webhook).To(Equal("http://test.foo.com"), "webhook should be http://test.foo.com")
		})

	})

	Describe("UpdateWebhook", func() {
		It("returns item with new webhook", func() {
			rand.Seed(time.Now().UTC().UnixNano())
			url := fmt.Sprintf("http://plaid.com/webhook/%d", rand.Intn(100))
			res, err := client.UpdateWebhook(test.UpdateWebhookAccessToken, url)
			Expect(err).To(BeNil(), "err should be nil")
			Expect(res.Item).ToNot(BeNil(), "item should not be nil")
			Expect(res.Item.Webhook).To(Equal(url), "webhook should be"+url)
		})

	})
})
