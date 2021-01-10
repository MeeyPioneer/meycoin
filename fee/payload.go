package fee

import (
	"math/big"
)

const (
	baseTxFee            = "2000000000000000" // 0.002 MEYCOIN
	payloadMaxSize       = 200 * 1024
	StateDbMaxUpdateSize = payloadMaxSize
	freeByteSize         = 200
)

var (
	baseTxMeyCoin   *big.Int
	zeroFee       bool
	stateDbMaxFee *big.Int
	aerPerByte    *big.Int
)

func init() {
	baseTxMeyCoin, _ = new(big.Int).SetString(baseTxFee, 10)
	zeroFee = false
	aerPerByte = big.NewInt(5000000000000) // 5,000 GAER, feePerBytes * PayloadMaxBytes = 1 MEYCOIN
	stateDbMaxFee = new(big.Int).Mul(aerPerByte, big.NewInt(StateDbMaxUpdateSize-freeByteSize))
}

func EnableZeroFee() {
	zeroFee = true
}

func DisableZeroFee() {
	zeroFee = false
}

func IsZeroFee() bool {
	return zeroFee
}

func NewZeroFee() *big.Int {
	return big.NewInt(0)
}

func PayloadTxFee(payloadSize int) *big.Int {
	if IsZeroFee() {
		return NewZeroFee()
	}
	if payloadSize == 0 {
		return new(big.Int).Set(baseTxMeyCoin)
	}
	size := paymentDataSize(int64(payloadSize))
	if size > payloadMaxSize {
		size = payloadMaxSize
	}
	return new(big.Int).Add(
		baseTxMeyCoin,
		new(big.Int).Mul(
			aerPerByte,
			big.NewInt(size),
		),
	)
}

func MaxPayloadTxFee(payloadSize int) *big.Int {
	if IsZeroFee() {
		return NewZeroFee()
	}
	if payloadSize == 0 {
		return new(big.Int).Set(baseTxMeyCoin)
	}
	return new(big.Int).Add(PayloadTxFee(payloadSize), stateDbMaxFee)
}

func paymentDataSize(dataSize int64) int64 {
	pSize := dataSize - freeByteSize
	if pSize < 0 {
		pSize = 0
	}
	return pSize
}

func PaymentDataFee(dataSize int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(paymentDataSize(dataSize)), aerPerByte)
}
