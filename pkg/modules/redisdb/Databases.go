package redisdb

type Nodedatabase int

const (
	Bridge             Nodedatabase = 0
	UserDatabase       Nodedatabase = 1
	PermissionDatabase Nodedatabase = 1
	AddressDatabase    Nodedatabase = 1
	AccountDatabase    Nodedatabase = 1
	ProductDatabase    Nodedatabase = 2
	CategoryDatabase   Nodedatabase = 2
	ShopCartDatabase   Nodedatabase = 2
	FreightDatabase    Nodedatabase = 2
)
