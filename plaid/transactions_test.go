package plaid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		b, err := ioutil.ReadFile("test/test_transactions.txt")
		if err != nil {
			t.Fatal("Unable to load json response")
		}
		fmt.Fprintln(w, string(b))
	}))
	defer ts.Close()
	var testBed environmentURL = environmentURL(ts.URL)

	client := NewCustomClient("foo", "bar", testBed, ts.Client())

	response, err := client.GetTransactions("foo", time.Now().AddDate(-3, 0, 0), time.Now(), -1, -1)
	if err != nil {
		t.Fatalf("Unexpected error from GetTransactions %v", err)
	}
	assert.Equal(t, "ZVWXnQdo1xIJQ7BG3d5nu9qejdBrKnCbVRxoR", response.Item.ItemID, "Item ID should be ZVWXnQdo1xIJQ7BG3d5nu9qejdBrKnCbVRxoR")
	assert.Equal(t, 4, len(response.Accounts), "There should be 4 accounts")
	assert.Equal(t, 17, response.TotalTransactions, "Total Transactions count should be 17")
	assert.Equal(t, 17, len(response.Transactions), "There should be 17 transactions")
	assert.Equal(t, float64(25), response.Transactions[0].Amount, "The first transaction should have an amount of 25")
	assert.Equal(t, "A3E8zAqDGlsgARBPqNLdCXj7pvnVeLhyy6xEb", response.Transactions[0].TransactionID, "The first transaction should have an id of A3E8zAqDGlsgARBPqNLdCXj7pvnVeLhyy6xEb")
}
