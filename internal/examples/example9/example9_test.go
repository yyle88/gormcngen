package example9

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example9/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global database instance for association fields testing
// 用于关联字段测试的全局数据库实例
var caseDB *gorm.DB

// TestMain sets up database with association models
// Creates test data for testing association field handling and relationship queries
//
// TestMain 为关联模型设置数据库
// 创建测试数据来测试关联字段处理和关系查询
func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// Auto migrate both models with their association relationship
	// 自动迁移两个模型及其关联关系
	done.Done(db.AutoMigrate(&models.User{}, &models.Profile{}))

	// Create test users
	// 创建测试用户
	const userCount = 5
	users := make([]*models.User, 0, userCount)
	for idx := 0; idx < userCount; idx++ {
		users = append(users, &models.User{
			Name: "User" + strconv.Itoa(idx+1),
		})
	}
	done.Done(db.Create(&users).Error)

	// Create profiles for users (has_one relationship)
	// 为用户创建配置文件（has_one关系）
	profiles := make([]*models.Profile, 0, userCount)
	for idx, user := range users {
		profiles = append(profiles, &models.Profile{
			Bio:    "Bio for User" + strconv.Itoa(idx+1),
			UserID: user.ID,
		})
	}
	done.Done(db.Create(&profiles).Error)

	caseDB = db
	m.Run()
}

// TestExample9 demonstrates association field handling in column generation
// Validates that association fields are correctly skipped while maintaining proper column mappings
//
// TestExample9 演示列生成中的关联字段处理
// 验证关联字段被正确跳过同时保持正确的列映射
func TestExample9(t *testing.T) {
	// Test User model columns - should NOT include Profile field
	// 测试User模型列 - 不应包含Profile字段
	user := &models.User{}
	userColumns := user.Columns()
	t.Log("User columns:", neatjsons.S(userColumns))

	// Verify User columns contain expected fields but not association field
	// 验证User列包含预期字段但不包含关联字段
	userColumnsJSON := neatjsons.S(userColumns)
	require.Contains(t, userColumnsJSON, `"ID"`)
	require.Contains(t, userColumnsJSON, `"id"`)
	require.Contains(t, userColumnsJSON, `"Name"`)
	require.Contains(t, userColumnsJSON, `"name"`)
	// Profile field should NOT appear in column mappings
	// Profile字段不应出现在列映射中
	require.NotContains(t, userColumnsJSON, "Profile")

	// Test Profile model columns
	// 测试Profile模型列
	profile := &models.Profile{}
	profileColumns := profile.Columns()
	t.Log("Profile columns:", neatjsons.S(profileColumns))

	// Test basic query with User model
	// 测试User模型的基础查询
	testBasicUserQuery(t, caseDB)

	// Test association query (User with Profile)
	// 测试关联查询（User带Profile）
	testAssociationQuery(t, caseDB)
}

// testBasicUserQuery tests basic User model queries using generated columns
// Demonstrates that association fields don't interfere with normal column operations
//
// testBasicUserQuery 使用生成的列测试基础User模型查询
// 演示关联字段不会干扰正常的列操作
func testBasicUserQuery(t *testing.T, db *gorm.DB) {
	var users []*models.User
	require.NoError(t, db.Find(&users).Error)
	require.Greater(t, len(users), 0)
	t.Log("Found users:", neatjsons.S(users))

	// Verify we can use the generated column name for queries
	// 验证我们可以使用生成的列名进行查询
	user := &models.User{}
	userColumns := user.Columns()

	var foundUser models.User
	require.NoError(t, db.Where(userColumns.Name.Eq("User1")).First(&foundUser).Error)
	require.Equal(t, "User1", foundUser.Name)
	t.Log("Found user by name:", neatjsons.S(foundUser))

	// Verify the column name mapping is correct
	// 验证列名映射是否正确
	require.Equal(t, "name", string(userColumns.Name))
	require.Equal(t, "id", string(userColumns.ID))
}

// testAssociationQuery tests GORM association queries
// Demonstrates that association relationships work correctly even though fields are skipped in column generation
//
// testAssociationQuery 测试GORM关联查询
// 演示即使字段在列生成中被跳过，关联关系仍然正常工作
func testAssociationQuery(t *testing.T, db *gorm.DB) {
	// Query User with Profile preloaded
	// 查询User并预加载Profile
	var users []*models.User
	require.NoError(t, db.Preload("Profile").Find(&users).Error)
	require.Greater(t, len(users), 0)

	// Verify association loading worked
	// 验证关联加载正常工作
	for _, user := range users {
		if user.Profile != nil {
			require.NotEmpty(t, user.Profile.Bio)
			require.Equal(t, user.ID, user.Profile.UserID)
			t.Log("User with profile:", neatjsons.S(map[string]interface{}{
				"UserName":   user.Name,
				"ProfileBio": user.Profile.Bio,
			}))
		}
	}
}
