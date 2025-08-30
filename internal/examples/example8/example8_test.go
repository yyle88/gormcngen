package example8

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example8/internal/models"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global database instance for extended column operations testing
// 用于扩展列操作测试的全局数据库实例
var caseDB *gorm.DB

// TestMain sets up database with advanced column generation configuration
// Creates test data for testing extended column features and query operations
//
// TestMain 为高级列生成配置设置数据库
// 创建测试数据来测试扩展列特性和查询操作
func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
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

// TestExample validates consistency between traditional and enhanced column queries
// Ensures both query approaches produce identical results with extended features
//
// TestExample 验证传统和增强列查询的一致性
// 确保两种查询方法在扩展功能下产生相同的结果
func TestExample(t *testing.T) {
	// Execute both query methods and compare results
	// 执行两种查询方法并比较结果
	expected0Text := neatjsons.S(selectFunc0(t, caseDB))
	expected1Text := neatjsons.S(selectFunc1(t, caseDB))
	// Verify results are identical with enhanced configurations
	// 验证在增强配置下结果相同
	require.Equal(t, expected0Text, expected1Text)
}

// selectFunc0 demonstrates conventional raw SQL query method
// Baseline implementation using string literals for comparison
//
// selectFunc0 演示常规的原生 SQL 查询方法
// 使用字符串字面量的基准实现用于比较
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

// selectFunc1 demonstrates enhanced type-safe column query method
// Advanced implementation with extended column generation features
//
// selectFunc1 演示增强的类型安全列查询方法
// 带有扩展列生成特性的高级实现
func selectFunc1(t *testing.T, db *gorm.DB) []*UserOrder {
	// Initialize models and get their decorated columns
	// 初始化模型并获取其装饰列
	user := &models.User{}
	userColumns := user.TableColumns(gormcnm.NewTableDecoration(user.TableName()))
	order := &models.Order{}
	orderColumns := order.TableColumns(gormcnm.NewTableDecoration(order.TableName()))

	// Define result mapping structure
	// 定义结果映射结构
	userOrder := &UserOrder{}

	// Build enhanced query with extended column features
	// 使用扩展列特性构建增强查询
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
