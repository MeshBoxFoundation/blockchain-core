package key

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto"
)

func keyCreate() (privkey *ecdsa.PrivateKey, err error) {

	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	privKey := privateKeyECDSA
	//publicKey := privateKeyECDSA.PublicKey
	//addr := crypto.PubkeyToAddress(publicKey)

	/*	fmt.Println("**********************************************************************")
		fmt.Printf("privKey:[%s]\npublicKey:[%s]\naddress:[%s]\n", privKey, publicKey, addr)
		fmt.Println("**********************************************************************")*/

	return privKey, nil
}
