package calc

import (
	"fmt"
	"math"
	"strconv"
)

// Numby ...
type Numby struct {
	numerator   int64
	denominator int64
}

// NewFromNumber returns new Numby with number as the numerator
func NewFromNumber(number string) *Numby {
	numerator, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	return &Numby{numerator: numerator, denominator: 1}
}

// NewFromFraction returns a new Numby from numerator and denominator
func NewFromFraction(numerator, denominator string) *Numby {
	left, err := strconv.ParseInt(numerator, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	right, err := strconv.ParseInt(denominator, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	return &Numby{numerator: left, denominator: right}
}

// NewFromMixed returns a new Numby from whole, numerator and denominator, as an improper fraction
func NewFromMixed(whole, numerator, denominator string) *Numby {
	w, err := strconv.ParseInt(whole, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	left, err := strconv.ParseInt(numerator, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	right, err := strconv.ParseInt(denominator, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	return &Numby{numerator: (w * right) + left, denominator: right}
}

func (n *Numby) Add(other *Numby) *Numby {
	return &Numby{
		numerator:   (n.numerator * other.denominator) + (n.denominator * other.numerator),
		denominator: n.denominator * other.denominator,
	}
}

func (n *Numby) Sub(other *Numby) *Numby {
	return &Numby{
		numerator:   (n.numerator * other.denominator) - (n.denominator * other.numerator),
		denominator: n.denominator * other.denominator,
	}
}

func (n *Numby) Mul(other *Numby) *Numby {
	return &Numby{
		numerator:   n.numerator * other.numerator,
		denominator: n.denominator * other.denominator,
	}
}

func (n *Numby) Div(other *Numby) *Numby {
	return &Numby{
		numerator:   n.numerator * other.denominator,
		denominator: n.denominator * other.numerator,
	}
}

func (n *Numby) Simplify() string {
	if n.numerator == n.denominator {
		return fmt.Sprintf("1")
	}

	gcd := gcd(n.numerator, n.denominator)

	var nume, denom int64

	if 1 == math.Abs(float64(gcd)) {
		nume = n.numerator
		denom = n.denominator
	} else {
		nume = n.numerator / gcd
		denom = n.denominator / gcd
	}

	if nume < denom {
		return fmt.Sprintf("%d/%d", nume, denom)
	}

	if nume%denom == int64(0) {
		return fmt.Sprintf("%d", nume/denom)
	}

	return fmt.Sprintf("%d_%d/%d", nume/denom, nume%denom, denom)
}

func (n *Numby) String() string {
	return n.Simplify()
}

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
