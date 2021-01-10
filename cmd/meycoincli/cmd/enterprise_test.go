package cmd

import (
	"encoding/binary"
	"errors"
	"testing"

	"github.com/mr-tron/base58/base58"

	meycoinrpc "github.com/meeypioneer/meycoin/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetConfChangeWithMock(t *testing.T) {
	mock := initMock(t)
	defer deinitMock()

	var (
		testTxHashString    = "HB44gJvHhVoEfgiGq3VZmV9VUXfBXhHjcEvroBMkJGnY"
		testTxHash, _       = base58.Decode(testTxHashString)
		testBlockHashString = "56Qy6MQei9KM13rqEq1jiJ7Da21Kcq9KdmYWcnPLtxS3"
		testBlockHash, _    = base58.Decode(testBlockHashString)

		tx           *meycoinrpc.Tx        = &meycoinrpc.Tx{Hash: testTxHash, Body: &meycoinrpc.TxBody{Payload: []byte(string("{ \"name\": \"GetConfTest\" }"))}}
		resTxInBlock *meycoinrpc.TxInBlock = &meycoinrpc.TxInBlock{TxIdx: &meycoinrpc.TxIdx{BlockHash: testBlockHash, Idx: 1}, Tx: tx}

		expBlockNo = uint64(100)
	)

	// case: tx is not executed
	mock.EXPECT().GetTX(
		gomock.Any(), // expect any value for first parameter
		gomock.Any(), // expect any value for second parameter
	).Return(
		tx,
		nil,
	).AnyTimes()

	mock.EXPECT().GetReceipt(
		gomock.Any(), // expect any value for first parameter
		gomock.Any(), // expect any value for second parameter
	).Return(
		&meycoinrpc.Receipt{},
		nil,
	).MaxTimes(2)

	output, err := executeCommand(rootCmd, "enterprise", "tx", testTxHashString, "--timeout", "0")
	assert.NoError(t, err, "should be success")
	t.Log(output)

	// tx is executed
	blockNoBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(blockNoBytes, expBlockNo)

	mock.EXPECT().GetTX(
		gomock.Any(), // expect any value for first parameter
		gomock.Any(), // expect any value for second parameter
	).Return(
		nil,
		errors.New("tx is not in the main chain"),
	).MaxTimes(2)

	mock.EXPECT().GetBlockTX(
		gomock.Any(), // expect any value for first parameter
		gomock.Any(), // expect any value for second parameter
	).Return(
		resTxInBlock,
		nil,
	).MaxTimes(2)

	mock.EXPECT().GetBlock(
		gomock.Any(), // expect any value for first parameter
		gomock.Any(), // expect any value for second parameter
	).Return(
		&meycoinrpc.Block{Header: &meycoinrpc.BlockHeader{BlockNo: expBlockNo}},
		nil,
	).MaxTimes(2)

	state := meycoinrpc.ConfChangeState_CONF_CHANGE_STATE_APPLIED
	mock.EXPECT().GetConfChangeProgress(
		gomock.Any(), // expect any value for first parameter
		&meycoinrpc.SingleBytes{Value: blockNoBytes},
		//gomock.Any(), // expect any value for second parameter
	).Return(
		&meycoinrpc.ConfChangeProgress{State: state, Err: ""},
		nil,
	).MaxTimes(2)

	// case: GetConfChangeProgress from tx hash
	output, err = executeCommand(rootCmd, "enterprise", "tx", testTxHashString, "--timeout",  "0")
	assert.NoError(t, err, "should be success")
	t.Log(output)

	// case: GetConfChangeProgress from reqid
	_, err = executeCommand(rootCmd, "enterprise", "tx", testTxHashString, "--timeout", "0")
	assert.NoError(t, err, "should be success")
	t.Log(output)
}
