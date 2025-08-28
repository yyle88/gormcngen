package example7

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example7/internal/models"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global database instance for advanced column operations testing
// 用于高级列操作测试的全局数据库实例
var caseDB *gorm.DB

// TestMain sets up database with users and orders for join query testing
// Creates test data to demonstrate type-safe joins and complex queries
//
// TestMain 为联接查询测试设置用户和订单数据库
// 创建测试数据来演示类型安全的联接和复杂查询
func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := done.VCE(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&models.User{}, &models.Order{}))

	const userCount = 10
	users := make([]*models.User, 0, userCount)
	for idx := 0; idx < userCount; idx++ {
		users = append(users, &models.User{
			ID:   0,
			Name: "name" + strconv.Itoa(idx+1),
		})
	}
	done.Done(db.Create(&users).Error)

	const orderCount = 20
	orders := make([]*models.Order, 0, orderCount)
	for idx := 0; idx < orderCount; idx++ {
		userID := users[rand.IntN(len(users))].ID

		orders = append(orders, &models.Order{
			ID:     0,
			UserID: userID,
			Amount: float64(rand.IntN(1000)) + rand.Float64(),
		})
	}
	done.Done(db.Create(&orders).Error)

	caseDB = db
	m.Run()
}

// TestExample demonstrates comparison between raw SQL and type-safe column queries
// Compares traditional string-based queries with generated column-based queries
//
// TestExample 演示原生 SQL 和类型安全列查询的对比
// 比较传统基于字符串的查询和生成的基于列的查询
func TestExample(t *testing.T) {
	// Execute traditional query and type-safe query
	// 执行传统查询和类型安全查询
	expected0Text := neatjsons.S(selectFunc0(t, caseDB))
	expected1Text := neatjsons.S(selectFunc1(t, caseDB))
	// Ensure both results are identical
	// 确保两者结果相同
	require.Equal(t, expected0Text, expected1Text)
}

// selectFunc0 demonstrates traditional raw SQL query approach
// Uses string literals for table names, column names, and joins
//
// selectFunc0 演示传统的原生 SQL 查询方法
// 使用字符串字面量表示表名、列名和联接
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

// selectFunc1 demonstrates type-safe column-based query approach
// Uses generated column structs with Chinese column names and type safety
//
// selectFunc1 演示类型安全的基于列的查询方法
// 使用生成的列结构体和中文列名及类型安全
func selectFunc1(t *testing.T, db *gorm.DB) []*UserOrder {
	// Create model instances and get their table-decorated columns
	// 创建模型实例并获取带表装饰的列
	user := &models.User{}
	userColumns := user.TableColumns(gormcnm.NewTableDecoration(user.TableName()))
	order := &models.Order{}
	orderColumns := order.TableColumns(gormcnm.NewTableDecoration(order.TableName()))

	// Create result struct for column name mapping
	// 为列名映射创建结果结构体
	userOrder := &UserOrder{}

	// Build type-safe query with Chinese column names
	// 使用中文列名构建类型安全查询
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
