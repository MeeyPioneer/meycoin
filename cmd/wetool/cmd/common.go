/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package cmd

import (
	"fmt"
	"github.com/meeypioneer/meycoin/cmd/meycoincli/util/encoding/json"
	"github.com/meeypioneer/meycoin/types"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type WezenClient struct {
	types.WezenRPCServiceClient
	conn *grpc.ClientConn

}


func GetClient(serverAddr string, opts []grpc.DialOption) interface{} {
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil || conn == nil {
		fmt.Println(err)
		panic("connection failed")
	}

	connClient := &WezenClient{
		WezenRPCServiceClient: types.NewWezenRPCServiceClient(conn),
		conn: conn,
	}

	return connClient
}

func (c *WezenClient) Close() {
	c.conn.Close()
	c.conn = nil
}

// JSON converts protobuf message(struct) to json notation
func JSON(pb proto.Message) string {
	jsonout, err := json.MarshalIndent(pb, "", " ")
	if err != nil {
		fmt.Printf("Failed: %s\n", err.Error())
		return ""
	}
	return string(jsonout)
}
