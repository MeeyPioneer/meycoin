package cmd

import (
	"bytes"
	"testing"

	"github.com/meeypioneer/meycoin/cmd/meycoincli/cmd/mock_types"
	"github.com/meeypioneer/meycoin/cmd/meycoincli/util"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
)

func initMock(t *testing.T) *mock_types.MockMeyCoinRPCServiceClient {
	test = true
	ctrl := gomock.NewController(t)
	mock := mock_types.NewMockMeyCoinRPCServiceClient(ctrl)
	mockClient := &util.ConnClient{
		MeyCoinRPCServiceClient: mock,
	}
	client = mockClient
	return mock
}

func deinitMock() {
	test = false
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()

	return c, buf.String(), err
}
