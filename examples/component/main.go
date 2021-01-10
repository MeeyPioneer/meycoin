/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/meeypioneer/mey-library/log"
	"github.com/meeypioneer/meycoin/examples/component/message"
	"github.com/meeypioneer/meycoin/examples/component/server"
	"github.com/meeypioneer/meycoin/examples/component/service"
	"github.com/meeypioneer/meycoin/pkg/component"
)

func main() {

	compHub := component.NewComponentHub()

	testServer := &server.TestServer{}
	testServer.BaseComponent = component.NewBaseComponent("TestServer", testServer, log.Default())

	helloService := service.NexExampleServie("Den")

	compHub.Register(testServer)
	compHub.Register(helloService)
	compHub.Start()

	// request and go through
	testServer.RequestTo(message.HelloService, &message.HelloReq{Who: "Roger"})

	// request and wait
	rawResponse, err := compHub.RequestFuture(message.HelloService,
		&component.CompStatReq{SentTime: time.Now()},
		time.Second, "examples/component.main").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		response := rawResponse.(*component.CompStatRsp)
		fmt.Printf("RequestFuture Test Result: %v\n", response)
	}

	// collect all component's statuses
	statics, _ := compHub.Statistics(time.Second, "")
	if data, err := json.MarshalIndent(statics, "", "\t"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("All Component's Statistics: %s\n", data)
	}

	time.Sleep(1 * time.Second)

	compHub.Stop()
}
