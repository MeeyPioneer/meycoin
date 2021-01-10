package fee

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPayloadTxFee(t *testing.T) {
	type args struct {
		payloadSize int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			"zero",
			args{payloadSize: 0},
			baseTxMeyCoin,
		},
		{
			"under200",
			args{payloadSize: 198},
			baseTxMeyCoin,
		},
		{
			"exact200",
			args{payloadSize: 198},
			baseTxMeyCoin,
		},
		{
			"over200",
			args{payloadSize: 265},
			new(big.Int).Add(baseTxMeyCoin, new(big.Int).Mul(new(big.Int).SetUint64(65), aerPerByte)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PayloadTxFee(tt.args.payloadSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PayloadTxFee() = %v, want %v", got, tt.want)
			}
		})
	}
}
