package redisdb

type Nodedatabase int

const (
	Bridge          Nodedatabase = 0
	UserDatabase    Nodedatabase = 1
	AddressDatabase Nodedatabase = 1
	AccountDatabase Nodedatabase = 1
	ProductDatabase Nodedatabase = 2
)
