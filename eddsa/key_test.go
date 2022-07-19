package eddsa

//func TestPublicKeyGeneration(t *testing.T) {
//	//secretKey := bigint.StringToBigInt("9d61b19deffd5a60ba844af492ec2cc44449c5697b326919703bac031cae7f60", 16)
//	//secretKey := bigint.StringToBigInt("4449c5697b326919703bac031cae7f609d61b19deffd5a60ba844af492ec2cc4", 16)
//	//
//	//privateKey := NewPrivateKey(ed25519.GetEd25519(), secretKey)
//	//publicKey := privateKey.PublicKey
//
//	//fmt.Printf("x: %x\n", publicKey.GetX().GetNum().Bytes())
//	//fmt.Printf("y: %x\n", publicKey.GetY().GetNum().Bytes())
//
//	//fmt.Printf("%v\n", publicKey.Bytes())
//
//	//fmt.Println(bigint.SwapEndianness(publicKey))
//	//g := ed25519.GetEd25519().GetG()
//	//expectedPK := bigint.StringToBigInt("d75a980182b10ab7d54bfed3c964073a0ee172f3daa62325af021a68f707511a", 16)
//	//
//	//if publicKey.Equal(expectedPK) == false {
//	//	t.Fatal("the assertion has failed")
//	//}
//}
//
//func TestKeyGeneration(t *testing.T) {
//	sk, _ := new(big.Int).SetString("9d61b19deffd5a60ba844af492ec2cc44449c5697b326919703bac031cae7f60", 16)
//
//	sk = new(big.Int).SetBit(sk, 0, 0)
//	sk = new(big.Int).SetBit(sk, 1, 0)
//	sk = new(big.Int).SetBit(sk, 2, 0)
//
//	// setting the highest bit to 0
//	sk = new(big.Int).SetBit(sk, 255, 0)
//
//	// setting the second highest bit to 1
//	sk = new(big.Int).SetBit(sk, 254, 1)
//
//	fmt.Printf("%x\n", sk.Bytes())
//	defs := ed25519.GetEd25519()
//	pkPoint := defs.GetG().ScalarMul(sk)
//	pk := pkPoint.GetY().GetNum().Bytes()
//	//fmt.Println(pk)
//	changeEndianness(&pk)
//	//fmt.Println(pk)
//	pk[31] &= 127
//	aa := byte((pkPoint.GetX().GetNum().Bytes()[0] & 1) << 7)
//	aaa := byte((pkPoint.GetX().GetNum().Bit(7)) << 7)
//
//	fmt.Println(pkPoint.GetX().GetNum().Bit(7))
//
//	fmt.Println(aa)
//	fmt.Println(aaa)
//
//	pk[31] |= aaa
//
//	//pk := new(big.Int).Set(pkPoint.GetY().GetNum())
//	//
//	//pk = new(big.Int).SetBit(pk, 7, 0)
//	//xLSB := pkPoint.GetX().GetNum().Bit(0)
//	//pk = new(big.Int).SetBit(pk, 255, xLSB)
//	want := "1d87a9026fd0126a5736fe1628c95dd419172b5b618457e041c9c861b2494a94"
//	//
//	fmt.Printf("%x\n", pk)
//	fmt.Printf("%s\n", want)
//
//}
//
//func TestBlah(t *testing.T) {
//	//n := new(big.Int).SetInt64(5)
//	//fmt.Println(n.Bit(4))
//	//fmt.Println(BitCount(n))
//	//fmt.Println(n.SetBit(n, 0, 0))
//	//for i := 0; i < len(n.Bits()); i++ {
//	//	fmt.Println(n.Bit())
//	//}
//}
//
//func TestBlah1(t *testing.T) {
//	//sk, pk, err := ed255192.GenerateKey(rand.Reader)
//}
//
//func changeEndianness(x *[]uint8) { // need to test that
//	b := *x
//	for i := 0; i < len(b)/2; i++ {
//		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
//	}
//}
