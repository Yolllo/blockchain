package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ElrondNetwork/elrond-go/api/groups"
	"github.com/ElrondNetwork/elrond-go/integrationTests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionGroup(t *testing.T) {
	node := integrationTests.NewTestProcessorNodeWithTestWebServer(3, 0, 0)

	testTransactionGasCostWithMissingFields(t, node)
}

func testTransactionGasCostWithMissingFields(tb testing.TB, node *integrationTests.TestProcessorNodeWithTestWebServer) {
	// this is an example found in the wild, should not add more fields in order to pass the tests
	tx := groups.SendTxRequest{
		Sender:   "yol19mfd998dpqcdm5m5z5gn7u3lw6jwnxk8w75gh4gzjxcxfvtxmm8q3gd0su",
		Receiver: "yol1x9wak9qv3wzcmfamuql3xweeurktypmz6uteege7mpq6dwsf3xts0m9kkd",
		Value:    "100",
		Nonce:    0,
		GasPrice: 100,
		Version:  1,
		ChainID:  "T",
	}

	jsonBytes, _ := json.Marshal(tx)
	req, _ := http.NewRequest("POST", "/transaction/cost", bytes.NewBuffer(jsonBytes))

	resp := node.DoRequest(req)
	require.NotNil(tb, resp)

	type transactionCostResponseData struct {
		Cost uint64 `json:"txGasUnits"`
	}
	type transactionCostResponse struct {
		Data  transactionCostResponseData `json:"data"`
		Error string                      `json:"error"`
		Code  string                      `json:"code"`
	}

	txCost := &transactionCostResponse{}
	loadResponse(tb, resp.Body, txCost)
	assert.Empty(tb, txCost.Error)

	assert.Equal(tb, integrationTests.MinTxGasLimit, txCost.Data.Cost)
}

func loadResponse(tb testing.TB, rsp io.Reader, destination interface{}) {
	jsonParser := json.NewDecoder(rsp)
	err := jsonParser.Decode(destination)

	assert.Nil(tb, err)
}
