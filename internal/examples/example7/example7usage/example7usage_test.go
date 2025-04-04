package example7usage

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example7"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&example7.User{}, &example7.Order{}))

	const userCount = 10
	users := make([]*example7.User, 0, userCount)
	for idx := 0; idx < userCount; idx++ {
		users = append(users, &example7.User{
			ID:   0,
			Name: "name" + strconv.Itoa(idx+1),
		})
	}
	done.Done(db.Create(&users).Error)

	const orderCount = 20
	orders := make([]*example7.Order, 0, orderCount)
	for idx := 0; idx < orderCount; idx++ {
		userID := users[rand.IntN(len(users))].ID

		orders = append(orders, &example7.Order{
			ID:     0,
			UserID: userID,
			Amount: float64(rand.IntN(1000)) + rand.Float64(),
		})
	}
	done.Done(db.Create(&orders).Error)

	caseDB = db
	m.Run()
}

func TestExample(t *testing.T) {
	expected0Text := neatjsons.S(selectFunc0(t, caseDB))
	expected1Text := neatjsons.S(selectFunc1(t, caseDB))
	//确保两者结果相同
	require.Equal(t, expected0Text, expected1Text)
}

// 这是比较常规的逻辑
func selectFunc0(t *testing.T, db *gorm.DB) []*UserOrder {
	var results []*UserOrder
	require.NoError(t, db.Table("users").
		Select("users.id as user_id, users.name as user_name, orders.id as order_id, orders.amount as order_amount").
		Joins("left join orders on orders.user_id = users.id").
		Order("users.id asc, orders.id asc").
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

type UserOrder struct {
	UserID      uint
	UserName    string
	OrderID     uint
	OrderAmount float64
}

func selectFunc1(t *testing.T, db *gorm.DB) []*UserOrder {
	user := &example7.User{}
	userColumns := user.TableColumns(gormcnm.NewTableDecoration(user.TableName()))
	order := &example7.Order{}
	orderColumns := order.TableColumns(gormcnm.NewTableDecoration(order.TableName()))

	userOrder := &UserOrder{}

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.V主键.AsName(gormcnm.Cnm(userOrder.UserID, "user_id")),
			userColumns.V名字.AsName(gormcnm.Cnm(userOrder.UserName, "user_name")),
			orderColumns.V订单主键.AsName(gormcnm.Cnm(userOrder.OrderID, "order_id")),
			orderColumns.V订单金额.AsName(gormcnm.Cnm(userOrder.OrderAmount, "order_amount")),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.V用户主键.OnEq(userColumns.V主键))).
		Order(userColumns.V主键.Ob("asc").Ob(orderColumns.V订单主键.Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}
