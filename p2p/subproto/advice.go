/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package subproto

import (
	"github.com/meeypioneer/mey-library/log"
	"github.com/meeypioneer/meycoin/p2p/p2pcommon"
	"github.com/meeypioneer/meycoin/p2p/p2putil"
	"github.com/rs/zerolog"
	"time"
)

func WithTimeLog(handler p2pcommon.MessageHandler, logger *log.Logger, level zerolog.Level) p2pcommon.MessageHandler {
	handler.AddAdvice(&LogHandleTimeAdvice{logger: logger, level:level})
	return handler
}

type LogHandleTimeAdvice struct {
	logger    *log.Logger
	level     zerolog.Level
	timestamp time.Time
}

func (a *LogHandleTimeAdvice) PreHandle() {
	a.timestamp = time.Now()
}

func (a *LogHandleTimeAdvice) PostHandle(msg p2pcommon.Message, msgBody p2pcommon.MessageBody) {
	a.logger.WithLevel(a.level).
		Int64("elapsed", time.Since(a.timestamp).Nanoseconds()/1000).
		Str(p2putil.LogProtoID, msg.Subprotocol().String()).
		Str(p2putil.LogMsgID, msg.ID().String()).
		Msg("handle takes")
}
