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

func NewPrivateKey(c curv3.EcdsaCurve, d *big.Int) *PrivateKey {
	gP := c.GetG()
	d = new(big.Int).Mod(d, c.GetN())

	return &PrivateKey{
		Curve:     c,
		PublicKey: gP.Mul(d),
		Key:       d,
	}
}
