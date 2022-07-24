package eddsa

//type Signature struct {
//	R *point.Point
//	S *big.Int
//}

// good resources
// https://github.com/warner/python-pure25519/blob/master/pure25519/basic.py#L261
// https://ed25519.cr.yp.to/python/ed25519.py
// https://github.com/ZenGo-X/curv/blob/master/src/elliptic/curves/ed25519.rs
// https://sefiks.com/2018/12/24/a-gentle-introduction-to-edwards-curve-digital-signature-algorithm-eddsa/
// https://github.com/serengil

//func Sign(priv *PrivateKey, msgToSign string, curve curve.EdwardsCurve) *Signature {
//	//
//	////  r = hash(hash(privKey) + msg) mod q
//	//msgB := []byte(msgToSign)
//	////privHash := sha512.New().Sum(priv.Key.Bytes())
//	//privHash := sha512.New().Sum(msgB)
//	//
//	//privAndMsgHash := sha512.New().Sum(append(privHash, msgB...))
//	//r := new(big.Int).Mod(new(big.Int).SetBytes(privAndMsgHash), curve.GetP())
//	//
//	////   R = r * G
//	//R := curve.GetG().Mul(r)
//	//
//	////  h = hash(R + pubKey + msg) mod q
//	//preImage := append(R.GetX().GetNum().Bytes(), priv.PublicKey.GetX().GetNum().Bytes()...)
//	//preImage = append(preImage, msgB...)
//	//
//	//hDigest := sha512.New().Sum(preImage)
//	//h := new(big.Int).Mod(new(big.Int).SetBytes(hDigest), curve.GetP())
//	//
//	////  s = (r + h * privKey) mod q
//	//s := new(big.Int).Mod(bigint.Add(bigint.Mul(h, priv.Key), r), curve.GetP())
//	//
//	////  signature { R, s }
//	//return &Signature{S: s, R: R}
//}
//
//func Verify(publicKey *point.Point, msgToSign string, sig *Signature, curve curve.EdwardsCurve) bool {
//
//	msgB := []byte(msgToSign)
//
//	// h = hash(R + pubKey + msg) mod q
//	preImage := append(sig.R.GetX().GetNum().Bytes(), publicKey.GetX().GetNum().Bytes()...)
//	preImage = append(preImage, msgB...)
//	hDigest := sha512.New().Sum(preImage)
//	h := new(big.Int).Mod(new(big.Int).SetBytes(hDigest), curve.GetP())
//	//  P1 = s * G
//	p1 := curve.GetG().Mul(sig.S)
//
//	// P2 = R + h * pubKey
//	p2 := new(big.Int).Mod(bigint.Mul(h, publicKey.GetX().GetNum()), curve.GetP())
//	p2 = new(big.Int).Mod(bigint.Add(sig.R.GetX().GetNum(), p2), curve.GetP())
//
//	// P1 == P2
//	return p1.GetX().GetNum().Cmp(p2) == 0
//}
