package example8 // Package example8 我试试把注释写到这里行不行

type User struct {
	ID   uint   `cnm:"V主键"`
	Name string `cnm:"V名字"`
}

func (*User) TableName() string {
	return "users"
}

type Order struct {
	ID     uint    `cnm:"V订单主键"`
	UserID uint    `cnm:"V用户主键"`
	Amount float64 `cnm:"V订单金额"`
}

func (*Order) TableName() string {
	return "orders"
}
