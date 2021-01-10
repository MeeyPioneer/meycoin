/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package message

const HelloService = "HelloService"

type HelloReq struct{ Who string }

type HelloRsp struct{ Greeting string }
