package point

//
//// Curve ax² + y² = 1 − d ·x²·y²
//type Point struct {
//	a *big.Int
//	d *big.Int
//	x *field.Element
//	y *field.Element
//}
//
//func New(x, y *field.Element, a, d *big.Int) *Point {
//
//	//def isoncurve(P):
//	//x = P[0]
//	//y = P[1]
//	//return (-x*x + y*y - 1 - d*x*x*y*y) % Q == 0
//
//	res := &Point{
//		a: a,
//		d: d,
//		x: x,
//		y: y,
//	}
//
//	return res
//}
//
//func (p *Point) Add(p2 *Point) *Point {
//	if p.a.Cmp(p2.a) != 0 || p.d.Cmp(p2.d) != 0 {
//		panic("points are not on the same curve")
//	}
//
//	// x_1 * y_2 + y_1 * x_2			==> a
//	a_1 := p.x.Mul(p2.y)
//	a_2 := p.y.Mul(p2.x)
//	a := a_1.Add(a_2)
//
//	// 1 + d * x_1 * y_1 * x_2 * y_2	==> b
//	b := new(big.Int).Mul(p.d, p.x.GetNum())
//	b = new(big.Int).Mul(b, p.y.GetNum())
//	b = new(big.Int).Mul(b, p2.x.GetNum())
//	b = new(big.Int).Mul(b, p2.y.GetNum())
//
//	one := new(big.Int).SetInt64(1)
//	b = new(big.Int).Add(b, one)
//
//	// y_1 * y_2 - a * x_1 * x_2		==> C
//	cLeft := p.y.Mul(p2.y)
//	cRight := new(big.Int).Mul(new(big.Int).Mul(p.a, p.x.GetNum()), p2.x.GetNum())
//	cRight = new(big.Int).Mod(cRight, p.x.GetOrder())
//	c := new(big.Int).Sub(cLeft.GetNum(), cRight)
//
//	// 1 - d * x_1 * y_1 * x_2 * y_2 	==> D
//	dRight := new(big.Int).Mul(p.d, new(big.Int).Mul(p.x.GetNum(), new(big.Int).Mul(p.y.GetNum(), new(big.Int).Mul(p2.x.GetNum(), p2.y.GetNum()))))
//	dRight = new(big.Int).Mod(dRight, p.x.GetOrder())
//	d := new(big.Int).Sub(one, dRight)
//
//	// a / b = x_3
//	bInverse := new(big.Int).ModInverse(b, p.x.GetOrder())
//	x := a.Mul(bInverse)
//
//	// c / d = y_3
//	dInverse := new(big.Int).ModInverse(d, p.x.GetOrder())
//	y := field.New(new(big.Int).Mod(new(big.Int).Mul(c, dInverse), p.x.GetOrder()), p.x.GetOrder())
//
//	return &Point{
//		x: x,
//		y: y,
//		a: new(big.Int).Set(p.a),
//		d: new(big.Int).Set(p.d),
//	}
//}
//
//func (p *Point) Mul(coef *big.Int) *Point {
//	a := new(big.Int).Set(p.a)
//	d := new(big.Int).Set(p.d)
//
//	curr := Clone(p)
//	res := New(field.New(new(big.Int).SetInt64(0), p.x.GetOrder()), field.New(new(big.Int).SetInt64(1), p.x.GetOrder()), a, d)
//
//	for coef.Cmp(big.NewInt(0)) > 0 {
//		if coef.Bit(0) == 1 {
//			res = res.Add(curr)
//		}
//		curr = curr.Add(curr)
//		coef = new(big.Int).Rsh(coef, 1)
//	}
//
//	return res
//}
//
//func (p *Point) Equal(p2 *Point) bool {
//	if p.x == nil || p2.x == nil || p.y == nil || p2.y == nil {
//		if p.x != nil || p2.x != nil || p.y != nil || p2.y != nil {
//			panic("equal function validation error: invalid infinity point definition")
//		}
//
//		return p.d.Cmp(p2.d) == 0 && p.a.Cmp(p2.a) == 0
//	}
//
//	return p.d.Cmp(p2.d) == 0 && p.x.Equal(p2.x) && p.y.Equal(p2.y) && p.a.Cmp(p2.a) == 0
//}
//
//func Clone(p *Point) *Point {
//	newA := new(big.Int).Set(p.a)
//	newD := new(big.Int).Set(p.d)
//	newX := field.Clone(p.x)
//	newY := field.Clone(p.y)
//
//	return &Point{a: newA, d: newD, x: newX, y: newY}
//}
//
//func (p *Point) Print() string {
//	res := "{x: "
//	if p.x == nil {
//		res += "nil, "
//	} else {
//		res += p.x.ToString() + ", "
//	}
//
//	res += "y: "
//	if p.y == nil {
//		res += "nil"
//	} else {
//		res += p.y.ToString() + ", "
//	}
//
//	res += "a: "
//	if p.a == nil {
//		res += "nil"
//	} else {
//		res += p.a.String() + ", "
//	}
//
//	res += "d: "
//
//	if p.d == nil {
//		res += "nil"
//	} else {
//		res += p.d.String() + ", "
//	}
//
//	res += "}"
//
//	return res
//}
//
//func (p *Point) GetX() *field.Element {
//	return p.x
//}
//
//func (p *Point) GetY() *field.Element {
//	return p.y
//}
