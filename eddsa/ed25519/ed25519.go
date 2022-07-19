package ed25519

//var Ed25519Defs curv3.EdwardsCurve
//
//func GetEd25519() curv3.EdwardsCurve {
//	if Ed25519Defs == nil {
//		Ed25519Defs = newEd25519()
//	}
//
//	return Ed25519Defs
//}
//
//func newEd25519() *Ed25519 {
//	curveOrder := new(big.Int).Exp(new(big.Int).SetInt64(2), new(big.Int).SetInt64(252), nil)
//	toAdd, _ := new(big.Int).SetString("27742317777372353535851937790883648493", 10)
//	curveOrder = new(big.Int).Add(curveOrder, toAdd)
//
//	p := new(big.Int).Sub(new(big.Int).Exp(new(big.Int).SetInt64(2), new(big.Int).SetInt64(255), nil), new(big.Int).SetInt64(19))
//
//	a := big.NewInt(-1)
//
//	d := new(big.Int).SetInt64(-121665)
//	d = new(big.Int).Mod(new(big.Int).Mul(d, new(big.Int).ModInverse(new(big.Int).SetInt64(121666), p)), p)
//
//	gX, _ := new(big.Int).SetString("15112221349535400772501151409588531511454012693041857206046113283949847762202", 10)
//	gY, _ := new(big.Int).SetString("46316835694926478169428394003475163141307993866256225615783033603165251855960", 10)
//
//	gP := point.New(field.New(gX, p), field.New(gY, p), a, d)
//
//	return &Ed25519{
//		n:      curveOrder,
//		p:      p,
//		d:      d,
//		a:      a,
//		gPoint: gP,
//	}
//}
//
//type Ed25519 struct {
//	n      *big.Int
//	p      *big.Int
//	d      *big.Int
//	a      *big.Int
//	gPoint *point.Point
//}
//
//func (s *Ed25519) GetG() *point.Point {
//	return s.gPoint
//}
//
//func (s *Ed25519) GetD() *big.Int {
//	return s.d
//}
//
//func (s *Ed25519) GetA() *big.Int {
//	return s.a
//}
//
//func (s *Ed25519) GetP() *big.Int {
//	return s.p
//}
//
//func (s *Ed25519) GetN() *big.Int {
//	return s.n
//}
