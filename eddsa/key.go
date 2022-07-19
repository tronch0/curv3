package eddsa

//
//// key construction reference - https://raw.githubusercontent.com/LedgerHQ/orakolo/master/papers/Ed25519_BIP%20Final.pdf
//type PrivateKey struct {
//	Curve      curv3.EdwardsCurve
//	PrivateKey *big.Int
//	//Nonce     *big.Int
//	PublicKey *big.Int
//}
//
//func NewPrivateKey(defs curv3.EdwardsCurve, k *big.Int) *PrivateKey {
//	sk, pk := generateKeyPair(k, defs)
//
//	return &PrivateKey{
//		PrivateKey: sk,
//		Curve:      defs,
//		PublicKey:  pk,
//		//Nonce:     nonce,
//	}
//}
//
//func generateKeyPair(k *big.Int, defs curv3.EdwardsCurve) (*big.Int, *big.Int) {
//
//	sk := kToPrivateKey(k)
//	pk := privateToPublic(sk, defs)
//
//	return sk, pk
//
//	//var zInv, x, y field.Element
//	//zInv.Invert(&v.z)       // zInv = 1 / Z
//	//x.Multiply(&v.x, &zInv) // x = X / Z
//	//y.Multiply(&v.y, &zInv) // y = Y / Z
//	//
//	//out := copyFieldElement(buf, &y)
//	//out[31] |= byte(x.IsNegative() << 7)
//
//	//nonce := new(big.Int).SetBytes(h[32:])
//}
//
//func kToPrivateKey(k *big.Int) *big.Int {
//	h := sha512.New().Sum(k.Bytes())
//
//	// get the lower 32bytes and create d scalar out of it
//	sk := new(big.Int).SetBytes(h[:32])
//
//	// setting the lower 3 bits to 0
//	sk = new(big.Int).SetBit(sk, 0, 0)
//	sk = new(big.Int).SetBit(sk, 1, 0)
//	sk = new(big.Int).SetBit(sk, 2, 0)
//
//	// setting the highest bit to 0
//	sk = new(big.Int).SetBit(sk, 255, 0)
//
//	// setting the second highest bit to 1
//	sk = new(big.Int).SetBit(sk, 254, 1)
//	return sk
//}
//
//func privateToPublic(priv *big.Int, defs curv3.EdwardsCurve) *big.Int {
//	// treat it as d scalar and multipli with G
//	pkPoint := defs.GetG().ScalarMul(priv)
//
//	// take Y coor of the public point
//	pk := new(big.Int).Set(pkPoint.GetY().GetNum())
//
//	// set the most significant bit to 0
//	pk = new(big.Int).SetBit(pk, 255, 0)
//
//	//To form the encoding of the point, copy the least
//	//significant bit of the x-coordinate to the most significant bit of
//	//the final octet.
//
//	// get the least significant bit of x
//	xLSB := pkPoint.GetY().GetNum().Bit(0)
//	pk = new(big.Int).SetBit(pk, 7, xLSB)
//
//	return pk
//}
//
////func (priv *PrivateKey) GetPublicKeyEncoding() {
////	yB := priv.PublicKey.GetY().GetNum().Bytes()
////
////	changeEndianness(&yB)
////
////}
////
////func changeEndianness(x *[]uint8) { // need to test that
////	b := *x
////	for i := 0; i < len(b)/2; i++ {
////		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
////	}
////}
