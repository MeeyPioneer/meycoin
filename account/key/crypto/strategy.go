/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package key

import (
	"github.com/btcsuite/btcd/btcec"
)

type PrivateKey = btcec.PrivateKey

type KeyCryptoStrategy interface {
	Encrypt(key *PrivateKey, passphrase string) ([]byte, error)
	Decrypt(encrypted []byte, passphrase string) (*PrivateKey, error)
}
