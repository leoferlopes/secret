package crypto

import (
	"math"
)

func bytesOf(n byte) []byte {
	selectors := []byte{
		127,
		191,
		239,
		223,
		247,
		251,
		253,
		254,
	}

	var bytes []byte
	for i := 0; i < len(selectors); i++ {
		if n &^ selectors[i] != 0 {
			bytes = append(bytes, 1)
		} else {
			bytes = append(bytes, 0)
		}
	}
	return bytes
}

func bytesOfArray(b []byte) []byte {
	y := make([]byte, 0)
	for _, x := range b {
		y = append(y, bytesOf(x)...)
	}
	return y
}

func prepend(n []byte, size int) []byte {
	b := append([]byte{}, n...)
	b = append(b, make([]byte, size)...)
	return b
}

func xor(a []byte, b[]byte) []byte {
	c := make([]byte, len(a))
	b = prepend(b, len(a) - len(b))
	for i, x := range a {
		c[i] = x ^ b[i]
	}
	return c
}

func indexOf(a []byte, y byte) int {
	for i, x := range a {
		if x == y {
			return i
		}
	}
	return -1
}

func ltrim(a []byte) []byte {
	for i, x := range a {
		if x == 1 {
			return a[i:]
		}
	}
	return nil
}

func rshift(a []byte, size int) []byte {
	b := append(make([]byte, size), a...)
	return b
}

func binaryToInt(a []byte) int {
	x := 0
	for i, y := range a {
		x += int(y) * int(math.Pow(float64((2)), float64(len(a) - 1 - i)))
	}
	return x
}

func crc(n []byte) int {
	n = bytesOfArray(n)
	n = prepend(n, 7)
	k := bytesOf(13)
	for i := indexOf(k, 1); i < len(n) - 7 && i != -1; i = indexOf(k, 1) {
		k = ltrim(k)
		s := indexOf(n, 1)
		if s == -1 {
			break
		}
		k = rshift(k, s)
		n = xor(n, k)
	}
	return binaryToInt(n[8:])
}