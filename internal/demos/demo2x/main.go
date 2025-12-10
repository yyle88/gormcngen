// Package main: Advanced demo application showcasing gormcngen with complex queries
// Demonstrates table joins and sophisticated column operations using type-safe methods
// Shows database scenarios with account-purchase relationships in production settings
//
// main: 展示 gormcngen 复杂查询的高级演示应用程序
// 演示使用类型安全方法进行表连接和复杂的列操作
// 展示账户-采购关系的数据库场景
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/gormcngen/internal/demos/demo2x/internal/models"
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
	// Create new db connection // 创建数据库连接
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))

	// Create example data // 创建示例数据
	must.Done(db.AutoMigrate(&models.Account{}))
	must.Done(db.AutoMigrate(&models.Purchase{}))
	must.Done(db.Save(&models.Account{ID: 0, Name: "abc"}).Error)
	must.Done(db.Save(&models.Account{ID: 0, Name: "uvw"}).Error)
	must.Done(db.Save(&models.Account{ID: 0, Name: "xyz"}).Error)

	{
		var accounts []*models.Account
		must.Done(db.Find(&accounts).Error)
		zaplog.SUG.Debug(neatjsons.S(accounts))
	}

	must.Done(db.Save([]*models.Purchase{
		{ID: 0, AccountID: 1, ProductName: "A", Amount: 10},
		{ID: 0, AccountID: 1, ProductName: "B", Amount: 20},
		{ID: 0, AccountID: 1, ProductName: "C", Amount: 30},
		{ID: 0, AccountID: 2, ProductName: "U", Amount: 40},
		{ID: 0, AccountID: 2, ProductName: "V", Amount: 50},
		{ID: 0, AccountID: 2, ProductName: "W", Amount: 60},
		{ID: 0, AccountID: 3, ProductName: "X", Amount: 70},
		{ID: 0, AccountID: 3, ProductName: "Y", Amount: 80},
		{ID: 0, AccountID: 3, ProductName: "Z", Amount: 90},
	}).Error)

	{
		var purchases []*models.Purchase
		must.Done(db.Find(&purchases).Error)
		zaplog.SUG.Debug(neatjsons.S(purchases))
	}

	type AccountPurchase struct {
		AccountID      string  `gorm:"column:account_id;"`
		AccountName    string  `gorm:"column:account_name;"`
		PurchaseID     uint    `gorm:"column:purchase_id;"`
		ProductName    string  `gorm:"column:product_name;"`
		PurchaseAmount float64 `gorm:"column:purchase_amount;"`
	}

	{
		var results []*AccountPurchase
		must.Done(db.Table("accounts").
			Select("accounts.id as account_id, accounts.name as account_name, purchases.id as purchase_id, purchases.product_name, purchases.amount as purchase_amount").
			Joins("left join purchases on purchases.account_id = accounts.id").
			Order("accounts.id asc, purchases.id desc").
			Scan(&results).Error)
		/* Executed SQL // 执行的 SQL:
		SELECT accounts.id as account_id, accounts.name as account_name,
			   purchases.id as purchase_id, purchases.product_name,
			   purchases.amount as purchase_amount
		FROM `accounts`
		left join purchases on purchases.account_id = accounts.id
		ORDER BY accounts.id asc, purchases.id desc
		*/
		zaplog.SUG.Debug(neatjsons.S(results))
	}
	{
		account := &models.Account{}
		accountColumns := account.Columns()
		purchase := &models.Purchase{}
		purchaseColumns := purchase.Columns()

		accountPurchase := &AccountPurchase{}

		var results []*AccountPurchase
		must.Done(db.Table(account.TableName()).
			Select(gormcnmstub.MergeStmts(
				accountColumns.ID.WithTable(account).
					AsAlias("account_id"), // Use alias name // 直接使用别名
				accountColumns.Name.WithTable(account).
					AsName("account_name"), // Set target column name // 指定目标列名
				purchaseColumns.ProductName.WithTable(purchase).Name(), // Not recommend without alias // 不建议不指定别名
				purchaseColumns.ID.WithTable(purchase).
					AsName(gormcnm.Cnm(accountPurchase.PurchaseID, "purchase_id")), // Set target column with type protection // 指定目标列名避免类型不匹配
				purchaseColumns.Amount.WithTable(purchase).
					AsName(gormcnm.New[float64]("purchase_amount")), // Set target column with type constraint // 指定目标列名同时限制类型
			)).
			Joins(accountColumns.LEFTJOIN(purchase.TableName()).
				On(purchaseColumns.AccountID.WithTable(purchase).
					Eq(accountColumns.ID.WithTable(account)))).
			Order(accountColumns.ID.WithTable(account).Ob("asc").
				Ob(purchaseColumns.ID.WithTable(purchase).Ob("desc")).Ox()).
			Scan(&results).Error)
		/* Executed SQL // 执行的 SQL:
		SELECT accounts.id as account_id, accounts.name as account_name,
			   purchases.product_name, purchases.id as purchase_id,
			   purchases.amount as purchase_amount
		FROM `accounts`
		LEFT JOIN purchases ON purchases.account_id = accounts.id
		ORDER BY accounts.id asc , purchases.id desc
		*/
		zaplog.SUG.Debug(neatjsons.S(results))
	}
}
