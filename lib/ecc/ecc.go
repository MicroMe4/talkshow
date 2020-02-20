package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
)

//ECCRunner 用来精心
type ECCRunner struct {
	private *ecdsa.PrivateKey
	public  ecdsa.PublicKey
	cur     elliptic.Curve
}
