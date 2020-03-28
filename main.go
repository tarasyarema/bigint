package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Inspiration: https://en.wikipedia.org/wiki/Multiplication_algorithm#Optimizing_space_complexity

// BigInt is the generic BigInt struct
type BigInt struct {
	// the array data
	d []int
	// the base
	b int
}

// Multiply multiplies to BigInts
func (x BigInt) Multiply(y BigInt) (BigInt, error) {
	if x.b != y.b {
		return BigInt{}, fmt.Errorf("bases of x and y are not the same: %d, %d", x.b, y.b)
	}

	total := 0

	p, q := len(x.d), len(y.d)
	base := x.b

	prod := make([]int, p+q)

	for i := 0; i < p+q-1; i++ {
		jMin := Max(0, i-p+1)
		jMax := Min(i, q-1)

		for j := jMin; j <= jMax; j++ {
			total += (x.d[i-j] * y.d[j])
		}

		prod[i] = total % base
		total /= base
	}

	prod[p+q-1] = total % base

	return BigInt{prod, base}, nil
}

// String returns the string version of the number
func (x BigInt) String() string {
	s := ""

	for i := len(x.d) - 1; i >= 0; i-- {
		s += strconv.Itoa(x.d[i])
	}

	return strings.TrimLeftFunc(s, func(r rune) bool {
		if r == '0' {
			return true
		}

		return false
	})
}

// InitBigInt inits a BigInt from a string
func InitBigInt(s string, b int) (BigInt, error) {
	if b <= 1 {
		return BigInt{}, fmt.Errorf("base is not positive: %v", b)
	}

	a := make([]int, len(s))

	for i := 0; i < len(a); i++ {
		x, err := strconv.Atoi(fmt.Sprintf("%c", s[i]))

		if err != nil {
			return BigInt{}, err
		}

		a[len(a)-i-1] = x
	}

	return BigInt{a, b}, nil
}

func factorial(n int) (BigInt, error) {
	base := 10
	x := BigInt{[]int{1}, base}

	for i := 2; i <= n; i++ {
		y, err := InitBigInt(strconv.Itoa(i), base)
		if err != nil {
			return BigInt{}, err
		}

		z, err := x.Multiply(y)
		if err != nil {
			return BigInt{}, err
		}

		x = z
	}

	return x, nil
}

func main() {
	n := 1000

	t := time.Now()
	f, err := factorial(n)
	elapsed := time.Since(t)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d! = %v\n", n, f)
	fmt.Printf("elapsed: %v\n", elapsed)
}
