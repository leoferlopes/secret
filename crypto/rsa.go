package crypto

import "math/big"

type RSA struct {
	PublicKey     []byte
	SecretKey     []byte
	componentSize int
}

func (r *RSA) N() []byte {
	return r.PublicKey[:r.componentSize]
}

func (r *RSA) E() []byte {
	return r.PublicKey[r.componentSize:]
}

func (r *RSA) D() []byte {
	return r.SecretKey[r.componentSize:]
}

func NewRSA(size int) *RSA {
	p := randomPrime()
	q := randomPrime()

	for number().Sub(p, q).Uint64() == 0 {
		q = randomPrime()
	}
	n := number().Mul(q, p)
	z := number().Mul(number().Sub(p, number(1)), number().Sub(q, number(1)))
	e := randomPrime()
	for number().Sub(e, n).Sign() != -1 {
		e = randomPrime()
	}
	d := randomPrime()
	for true {
		ed := number()
		ed = number().Mul(e, d)
		m := number().Mod(ed, z)
		if m.Uint64() == 1 {
			break
		}
		d = d.Add(d, big.NewInt(1))
	}
	nb := padding(n.Bytes(), size)
	db := padding(d.Bytes(), size)
	eb := padding(e.Bytes(), size)
	return &RSA{
		PublicKey: append(nb, eb...),
		SecretKey: append(padding(n.Bytes(), 2), db...),
		componentSize: size,
	}
}