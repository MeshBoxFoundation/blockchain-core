package v1

type entry_v1 struct {
	nonce int32
	balance int32
}

func Newentry_v1(nonce int32, balance int32) *entry_v1 {
	r := &entry_v1{
		nonce:   nonce,
		balance: balance,
	}

	return r
}

func (ev1 *entry_v1) GetNonce() int32 {
	return ev1.nonce
}

func (ev1 *entry_v1) GetBalance() int32 {
	return ev1.balance
}