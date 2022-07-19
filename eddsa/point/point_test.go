package point

//	To test:
//	- New()
//	- Add()
//	- Equal()
//	- Clone()
//	- New()

//func TestScalarMul(t *testing.T) {
//	p := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(255), big.NewInt(0)), big.NewInt(19))
//	a := new(big.Int).SetInt64(-1)
//
//	curveOrder := new(big.Int).Exp(new(big.Int).SetInt64(2), new(big.Int).SetInt64(252), nil)
//	toAdd, _ := new(big.Int).SetString("27742317777372353535851937790883648493", 10)
//	curveOrder = new(big.Int).Add(curveOrder, toAdd)
//
//	d := new(big.Int).SetInt64(-121665)
//	d = new(big.Int).Mod(new(big.Int).Mul(d, new(big.Int).ModInverse(new(big.Int).SetInt64(121666), p)), p)
//
//	fmt.Printf("p: %d\n", p)                   // field order = 2^255 - 19 = 57896044618658097711785492504343953926634992332820282019728792003956564819949
//	fmt.Printf("curveOrder: %d\n", curveOrder) // curve order = 2^252 + 27742317777372353535851937790883648493 = 7237005577332262213973186563042994240857116359379907606001950938285454250989
//	fmt.Printf("ffA: %d\n", a)                 // - 1
//	fmt.Printf("ffD: %s\n", d)                 // 37095705934669439343138083508754565189542113879843219016388785533085940283555
//
//	gX, _ := new(big.Int).SetString("15112221349535400772501151409588531511454012693041857206046113283949847762202", 10)
//	gY, _ := new(big.Int).SetString("46316835694926478169428394003475163141307993866256225615783033603165251855960", 10)
//
//	testData := []struct {
//		coef      *big.Int
//		expectedX *big.Int
//		expectedY *big.Int
//	}{
//		{coef: big.NewInt(2), expectedX: stringToBigInt("24727413235106541002554574571675588834622768167397638456726423682521233608206", 10), expectedY: stringToBigInt("15549675580280190176352668710449542251549572066445060580507079593062643049417", 10)},
//		{coef: big.NewInt(5), expectedX: stringToBigInt("33467004535436536005251147249499675200073690106659565782908757308821616914995", 10), expectedY: stringToBigInt("43097193783671926753355113395909008640284023746042808659097434958891230611693", 10)},
//		{coef: bigint.SwapEndianness(stringToBigInt("12581e70a192aeb9ac1411b36d11fc06393db55998190491c063807a6b4d730d", 16)), expectedX: stringToBigInt("46953515626174660128743374276590207025464948126956050456964432034683890442435", 10), expectedY: stringToBigInt("43649996176441760651255662656482711906128939437336752974722489909985414406932", 10)},
//		{coef: bigint.SwapEndianness(stringToBigInt("0c2340b974bebfb9cb3f14e991bca432b57fb33f7c4d79e15f64209076afcd00", 16)), expectedX: stringToBigInt("48108495825706412711799803692360228025391948835486250305831184019146948949994", 10), expectedY: stringToBigInt("13228837014764440841117560545823854143168584625415590819123131242008409842892", 10)},
//	}
//
//	for i := range testData {
//		p1 := New(
//			field.New(gX, p), // x
//			field.New(gY, p), // y
//			a,                // a
//			d,                // d
//		)
//
//		var expectedRes *Point
//		expectedRes = New(
//			field.New(testData[i].expectedX, p), // x
//			field.New(testData[i].expectedY, p), // y
//			a,                                   // a
//			d,                                   // d
//		)
//
//		product := p1.ScalarMul(testData[i].coef)
//		if product.Equal(expectedRes) == false {
//			t.Fatal("assert result error")
//		}
//	}
//}
//
//func stringToBigInt(s string, base int) *big.Int {
//	res, _ := new(big.Int).SetString(s, base)
//	return res
//}
