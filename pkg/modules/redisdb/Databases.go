package redisdb

type Nodedatabase int

const (
	UserDatabase    Nodedatabase = 0
	Bridge          Nodedatabase = 1
	AddressDatabase Nodedatabase = 2
	AccountDatabase Nodedatabase = 3
)
