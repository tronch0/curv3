package ecdsa

import (
	"github.com/tronch0/curv3"
	"github.com/tronch0/curv3/ecdsa/point"
	"math/big"
)

type PrivateKey struct {
	Curve     curv3.EcdsaCurve `json:"curve"`
	Key       *big.Int         `json:"key"`
	PublicKey *point.Point     `json:"public_key"`
}

func NewPrivateKey(defs curv3.EcdsaCurve, k *big.Int) *PrivateKey {
	gP := defs.GetG()
	k = new(big.Int).Mod(k, defs.GetN())

	return &PrivateKey{
		Curve:     defs,
		PublicKey: gP.ScalarMul(k),
		Key:       k,
	}
}
