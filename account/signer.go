package account

import (
	"github.com/meeypioneer/mey-actor/actor"
	"github.com/meeypioneer/meycoin/account/key"
	"github.com/meeypioneer/meycoin/message"
)

type Signer struct {
	keystore *key.Store
}

func NewSigner(s *key.Store) *Signer {
	return &Signer{keystore: s}
}

//Receive actor message
func (s *Signer) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *message.SignTx:
		err := s.keystore.SignTx(msg.Tx, msg.Requester)
		defer context.Self().Stop()
		if err != nil {
			context.Respond(&message.SignTxRsp{Tx: nil, Err: err})
		} else {
			//context.Tell(context.Sender(), &message.SignTxRsp{Tx: msg, Err: nil})
			context.Respond(&message.SignTxRsp{Tx: msg.Tx, Err: nil})
		}
	}
}
