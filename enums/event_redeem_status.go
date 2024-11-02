package enums

type EventRedeemStatus uint64

const (
	NotRedeemed EventRedeemStatus = iota + 1
	Redeemed
)
