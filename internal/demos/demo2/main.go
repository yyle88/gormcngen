package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/gormcngen/internal/demos/demo2/internal/models"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//new db connection
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))

	//create example data
	must.Done(db.AutoMigrate(&models.User{}))
	must.Done(db.AutoMigrate(&models.Order{}))
	must.Done(db.Save(&models.User{ID: 0, Name: "abc"}).Error)
	must.Done(db.Save(&models.User{ID: 0, Name: "uvw"}).Error)
	must.Done(db.Save(&models.User{ID: 0, Name: "xyz"}).Error)

	{
		var users []*models.User
		must.Done(db.Find(&users).Error)
		zaplog.SUG.Debug(neatjsons.S(users))
	}

	must.Done(db.Save([]*models.Order{
		{ID: 0, UserID: 1, ProductName: "A", Amount: 10},
		{ID: 0, UserID: 1, ProductName: "B", Amount: 20},
		{ID: 0, UserID: 1, ProductName: "C", Amount: 30},
		{ID: 0, UserID: 2, ProductName: "U", Amount: 40},
		{ID: 0, UserID: 2, ProductName: "V", Amount: 50},
		{ID: 0, UserID: 2, ProductName: "W", Amount: 60},
		{ID: 0, UserID: 3, ProductName: "X", Amount: 70},
		{ID: 0, UserID: 3, ProductName: "Y", Amount: 80},
		{ID: 0, UserID: 3, ProductName: "Z", Amount: 90},
	}).Error)

	{
		var orders []*models.Order
		must.Done(db.Find(&orders).Error)
		zaplog.SUG.Debug(neatjsons.S(orders))
	}

	type UserOrder struct {
		UserID      string  `gorm:"column:user_id;"`
		UserName    string  `gorm:"column:user_name;"`
		OrderID     uint    `gorm:"column:order_id;"`
		ProductName string  `gorm:"column:product_name;"`
		OrderAmount float64 `gorm:"column:order_amount;"`
	}

	{
		var results []*UserOrder
		must.Done(db.Table("users").
			Select("users.id as user_id, users.name as user_name, orders.id as order_id, orders.product_name, orders.amount as order_amount").
			Joins("left join orders on orders.user_id = users.id").
			Order("users.id asc, orders.id desc").
			Scan(&results).Error)
		zaplog.SUG.Debug(neatjsons.S(results))
	}
	{
		user := &models.User{}
		userColumns := user.Columns()
		order := &models.Order{}
		orderColumns := order.Columns()

		userOrder := &UserOrder{}

		var results []*UserOrder
		must.Done(db.Table(user.TableName()).
			Select(gormcnmstub.MergeStmts(
				userColumns.ID.WithTable(user).
					AsAlias("user_id"), //直接使用别名
				userColumns.Name.WithTable(user).
					AsName("user_name"), //指定目标列名
				orderColumns.ProductName.WithTable(order).Name(), //这里不建议不指定别名
				orderColumns.ID.WithTable(order).
					AsName(gormcnm.Cnm(userOrder.OrderID, "order_id")), //指定目标列名，这是高级的用法能够避免类型不匹配
				orderColumns.Amount.WithTable(order).
					AsName(gormcnm.New[float64]("order_amount")), //指定目标列名，同时限制类型
			)).
			Joins(userColumns.LEFTJOIN(order.TableName()).
				On(orderColumns.UserID.WithTable(order).
					Eq(userColumns.ID.WithTable(user)))).
			Order(userColumns.ID.WithTable(user).Ob("asc").
				Ob(orderColumns.ID.WithTable(order).Ob("desc")).Ox()).
			Scan(&results).Error)
		zaplog.SUG.Debug(neatjsons.S(results))
	}
}
