/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package util

import (
	"fmt"

	"github.com/meeypioneer/meycoin/cmd/meycoincli/util/encoding/json"
	"github.com/meeypioneer/meycoin/types"
	protobuf "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type ConnClient struct {
	types.MeyCoinRPCServiceClient
	conn *grpc.ClientConn
}

func GetClient(serverAddr string, opts []grpc.DialOption) interface{} {
	conn := GetConn(serverAddr, opts)
	connClient := &ConnClient{
		MeyCoinRPCServiceClient: types.NewMeyCoinRPCServiceClient(conn),
		conn:                  conn,
	}

	return connClient
}

func GetConn(serverAddr string, opts []grpc.DialOption) *grpc.ClientConn {
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil || conn == nil {
		fmt.Println(err)
		panic("connection failed")
	}
	return conn
}

func (c *ConnClient) Close() {
	c.conn.Close()
	c.conn = nil
}

// JSON converts protobuf message(struct) to json notation
func JSON(pb protobuf.Message) string {
	jsonout, err := json.MarshalIndent(pb, "", " ")
	if err != nil {
		fmt.Printf("Failed: %s\n", err.Error())
		return ""
	}
	return string(jsonout)
}

func B58JSON(i interface{}) string {
	jsonout, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		fmt.Printf("Failed: %s\n", err.Error())
		return ""
	}
	return string(jsonout)
}
