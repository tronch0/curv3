package ecdsa

import (
	"github.com/tronch0/crypt0/bigint"
	"github.com/tronch0/curv3/ecdsa/secp256k1"
	"testing"
)

func TestPublicKeyGeneration(t *testing.T) {
	secretKey := bigint.GetRandom()
	privateKey := NewPrivateKey(secp256k1.GetSecp256k1(), secretKey)
	publicKey := privateKey.PublicKey

	g := secp256k1.GetSecp256k1().GetG()
	expectedPK := g.Mul(privateKey.Key)

	if publicKey.Equal(expectedPK) == false {
		t.Fatal("assertion has failed")
	}
}
