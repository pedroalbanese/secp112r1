// Parameters for the Secp112r1 Elliptic curve
package secp112r1

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var p112 *elliptic.CurveParams

func initP112() {
	p112 = new(elliptic.CurveParams)
	p112.P, _ = new(big.Int).SetString("db7c2abf62e35e668076bead208b", 16)
	p112.N, _ = new(big.Int).SetString("db7c2abf62e35e7628dfac6561c5", 16)
	p112.B, _ = new(big.Int).SetString("659ef8ba043916eede8911702b22", 16)
	p112.Gx, _ = new(big.Int).SetString("09487239995a5ee76b55f9c2f098", 16)
	p112.Gy, _ = new(big.Int).SetString("a89ce5af8724c0a23e0e0ff77500", 16)
	p112.BitSize = 128
	p112.Name = "secp128r1"
}

func P112() elliptic.Curve {
	initonce.Do(initP112)
	return p112
}
