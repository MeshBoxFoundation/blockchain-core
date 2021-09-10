package v1

type data_credits_entry_v1 struct {
	nonce int32
	balance int32
}

func Newdata_credits_entry_v1(nonce int32, balance int32) *data_credits_entry_v1 {
	r := &data_credits_entry_v1{
		nonce:   nonce,
		balance: balance,
	}

	return r
}

func (dce *data_credits_entry_v1) GetNonce() int32 {
	return dce.nonce
}

func (dce *data_credits_entry_v1) GetBalance() int32 {
	return dce.balance
}

