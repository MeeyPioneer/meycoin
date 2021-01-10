package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnDecodeAddress(t *testing.T) {
	addr, err := DecodeAddress("meycoin.system")
	assert.NoError(t, err, "decode meycoin.system")
	assert.Equal(t, "meycoin.system", string(addr), "decode meycoin.system")

	addr, err = DecodeAddress("meycoin.name")
	assert.NoError(t, err, "decode meycoin.name")
	assert.Equal(t, "meycoin.name", string(addr), "decode meycoin.name")

	addr, err = DecodeAddress("meycoin.enterprise")
	assert.NoError(t, err, "decode meycoin.enterprise")
	assert.Equal(t, "meycoin.enterprise", string(addr), "decode meycoin.enterprise")

	addr, err = DecodeAddress("")
	assert.NoError(t, err, "decode meycoin.name")
	assert.Equal(t, "", string(addr), "decode \"\"")

	encoded := EncodeAddress([]byte("meycoin.system"))
	assert.Equal(t, "meycoin.system", encoded, "encode meycoin.system")

	encoded = EncodeAddress([]byte("meycoin.name"))
	assert.Equal(t, "meycoin.name", encoded, "encode meycoin.name")

	encoded = EncodeAddress([]byte("meycoin.enterprise"))
	assert.Equal(t, "meycoin.enterprise", encoded, "encode meycoin.enterprise")

	encoded = EncodeAddress(nil)
	assert.Equal(t, "", encoded, "encode nil")

	testAddress := []byte{3, 49, 133, 46, 196, 51, 37, 27, 66, 1, 191, 92, 224, 12, 248, 246, 21, 247, 79, 10, 164, 63, 34, 97, 238, 0, 145, 129, 206, 54, 218, 241, 84}
	encoded = EncodeAddress(testAddress)
	assert.Equal(t, "AmNpn7K9wg6wsn6oMkTirQSUNdqtDm94iCrrpP5ZpwCAAxxPrsU2", encoded, "encode test address")

	addr, err = DecodeAddress("AmNpn7K9wg6wsn6oMkTirQSUNdqtDm94iCrrpP5ZpwCAAxxPrsU2")
	assert.NoError(t, err, "decode test address")
	assert.Equal(t, testAddress, addr, "decode test address")
}

func TestToAddress(t *testing.T) {
	addr := ToAddress("")
	assert.Equal(t, 0, len(addr), "nil")
}
