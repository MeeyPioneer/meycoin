/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package common

import (
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	"time"
)

// WezenMessage is data struct for transferring between wezen server and client.
// as of 2019.04.23, this is copy of MessageValue.
type WezenMessage struct {
	subProtocol p2pcommon.SubProtocol
	// Length is lenght of payload
	length uint32
	// timestamp is unix time (precision of second)
	timestamp int64
	// ID is 16 bytes unique identifier
	id p2pcommon.MsgID
	// OriginalID is message id of request which trigger this message. it will be all zero, if message is request or notice.
	originalID p2pcommon.MsgID

	// marshaled by google protocol buffer v3. object is determined by Subprotocol
	payload []byte
}


// NewWezenMessage create a new object
func NewWezenMessage(msgID p2pcommon.MsgID, protocol p2pcommon.SubProtocol, payload []byte) *WezenMessage {
	return &WezenMessage{id: msgID, timestamp:time.Now().UnixNano(), subProtocol:protocol,payload:payload,length:uint32(len(payload))}
}
func NewWezenRespMessage(msgID , orgReqID p2pcommon.MsgID, protocol p2pcommon.SubProtocol, payload []byte) *WezenMessage {
	return &WezenMessage{id: msgID, originalID:orgReqID, timestamp:time.Now().UnixNano(), subProtocol:protocol,payload:payload,length:uint32(len(payload))}
}

func (m *WezenMessage) Subprotocol() p2pcommon.SubProtocol {
	return m.subProtocol
}

func (m *WezenMessage) Length() uint32 {
	return m.length

}

func (m *WezenMessage) Timestamp() int64 {
	return m.timestamp
}

func (m *WezenMessage) ID() p2pcommon.MsgID {
	return m.id
}

func (m *WezenMessage) OriginalID() p2pcommon.MsgID {
	return m.originalID
}

func (m *WezenMessage) Payload() []byte {
	return m.payload
}

var _ p2pcommon.Message = (*WezenMessage)(nil)
