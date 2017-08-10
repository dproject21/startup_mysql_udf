package calcurate

import (
	"math/big"
)

func Calc(cmd string) string {
	switch cmd {
	case "pi":
		return calc_pi()
	default:
		return "no impliments."
	}
}

func calc_pi() string {
	max := int64(1000)
	two := big.NewRat(2, 1)
	pi := big.NewRat(1.0, 1.0)
	a := big.NewRat(1.0, 1.0)
	b := big.NewRat(1.0, 2.0)
	x := new(big.Rat)
	for i := int64(2); i < max; i += 2 {
		y := big.NewRat(i, i+1)
		a.Mul(a, y)
		x.Mul(a, b)
		pi.Add(pi, x)
		b.Quo(b, two) //割り算はQuoっぽい。第二引数が0だとpanic飛ばすらしいから、それが嫌なら予め篩うかInvで逆数を求めるか…
	}
	return pi.Mul(pi, two).FloatString(100)
}
