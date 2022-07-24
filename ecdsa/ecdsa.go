package ecdsa

import (
	"github.com/tronch0/crypt0/bigint"
	"github.com/tronch0/curv3"
	"github.com/tronch0/curv3/ecdsa/point"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func Sign(priv *PrivateKey, msgToSign string, c curv3.EcdsaCurve) *Signature {
	e := priv.Key
	z := bigint.HashStringToBigInt(msgToSign)
	k := bigint.GetRandomN(c.GetN())
	r := c.GetG().Mul(k).GetX().GetNum()
	kInverse := new(big.Int).ModInverse(k, c.GetN())

	s := new(big.Int).Mul(r, e)
	s = new(big.Int).Add(s, z)
	s = new(big.Int).Mul(s, kInverse)
	s = new(big.Int).Mod(s, c.GetN())

	return &Signature{R: r, S: s}
}

func Verify(publicKey *point.Point, z *big.Int, sig *Signature, c curv3.EcdsaCurve) bool {
	sigInverse := new(big.Int).ModInverse(sig.S, c.GetN())
	u := new(big.Int).Mul(z, sigInverse)
	u = new(big.Int).Mod(u, c.GetN())

	v := new(big.Int).Mul(sig.R, sigInverse)
	v = new(big.Int).Mod(v, c.GetN())

	uG := c.GetG().Mul(u)
	vP := publicKey.Mul(v)

	pointR := uG.Add(vP)
	return pointR.GetX().GetNum().Cmp(sig.R) == 0
}
