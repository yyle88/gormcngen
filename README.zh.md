[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcngen/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcngen/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcngen)](https://pkg.go.dev/github.com/yyle88/gormcngen)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcngen/master.svg)](https://coveralls.io/github/yyle88/gormcngen?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcngen.svg)](https://github.com/yyle88/gormcngen/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcngen)](https://goreportcard.com/report/github.com/yyle88/gormcngen)

# gormcngen

`gormcngen`: 赋予 GORM 模型使用 Columns() 获取列名的函数

就像 Java 生态系统中的 `MyBatis Plus`，它允许开发人员使用像 `Example::getName` 这样的表达式获取列名。

就像 Python 生态系统中的 `SQLAlchemy`，它允许开发人员使用像 `Example.name` 这样的 class 函数获得列名。

`gormcngen` 也为 Go 模型赋予 **类型安全** 的列引用功能。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

### 🔍 AST 级别精度
- **深度模型分析**: 解析结构体字段、标签和嵌入类型
- **GORM 标签提取**: 自动检测列名、类型和约束
- **嵌入字段支持**: 处理 `gorm.Model` 和自定义嵌入结构体
- **类型保护**: 在生成代码中维护精确的 Go 类型

### 🚀 智能代码生成
- **完美同步**: 生成的代码始终与你的模型匹配
- **自定义列名**: 遵循 `gorm:"column:name"` 标签
- **多语言支持**: 与 `cnm:"中文名"` 标签配合进行国际化开发
- **增量更新**: 只重新生成有变化的内容

### 🛠️ 开发体验
- **简单编程接口**: 易于使用的 Go API，立即获得结果
- **IDE 集成**: 生成的代码提供完整的智能提示支持
- **构建系统兼容**: 轻松集成 `go:generate` 指令
- **版本控制安全**: 确定性输出，确保清洁的差异

### 🏢 企业级就绪
- **大型代码库支持**: 轻松处理数百个模型
- **自定义命名约定**: 可配置的输出模式
- **验证和安全**: 内置检查防止无效生成
- **文档生成**: 自动生成的注释解释列映射

## 🏗️ 生态系统定位

```
┌─────────────────────────────────────────────────────────────────────┐
│                    GORM Type-Safe Ecosystem                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐              │
│  │  gormzhcn   │    │  gormmom    │    │  gormrepo   │              │
│  │ Chinese API │───▶│ Native Lang │───▶│  Package    │─────┐        │
│  │  Localize   │    │  Smart Tags │    │  Pattern    │     │        │
│  └─────────────┘    └─────────────┘    └─────────────┘     │        │
│         │                   │                              │        │
│         │                   ▼                              ▼        │
│         │            ┌─────────────┐              ┌─────────────┐   │
│         │            │ gormcngen   │              │Application  │   │
│         │            │Code Generate│─────────────▶│Custom Code  │   │
│         │            │AST Operation│              │             │   │
│         │            └─────────────┘              └─────────────┘   │
│         │                   │                              ▲        │
│         │                   ▼                              │        │
│         └────────────▶┌─────────────┐◄─────────────────────┘        │
│                       │   GORMCNM   │                               │
│                       │ FOUNDATION  │                               │
│                       │ Type-Safe   │                               │
│                       │ Core Logic  │                               │
│                       └─────────────┘                               │
│                              │                                      │
│                              ▼                                      │
│                       ┌─────────────┐                               │
│                       │    GORM     │                               │
│                       │  Database   │                               │
│                       └─────────────┘                               │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**gormcngen** 作为**代码生成引擎**，连接您的模型与类型安全基础层。

## 安装

```bash
go get github.com/yyle88/gormcngen
```

## 重要说明

**gormcngen 是一个 Go 包，不是 CLI 应用。** 它需要通过测试代码驱动的生成工作流程来使用。

## 快速开始

### 1. 创建项目结构

设置基本的项目结构，为模型和生成的代码创建专门的 DIR：

```bash
# 创建 models DIR
mkdir -p internal/models
```

### 2. 定义您的 GORM 模型

定义您的数据模型，gormcngen 将为这些模型生成列访问方法：

创建 `internal/models/models.go`：
```go
package models

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex;size:100"`
    Email    string `gorm:"index;size:255"`
    Age      int    `gorm:"column:age"`
    IsActive bool   `gorm:"default:true"`
}
```

### 3. 创建生成文件

创建用于存放生成代码的目标文件和包含生成逻辑的测试文件：

```bash
# 创建目标文件，包含 package 声明
echo "package models" > internal/models/ngen.go

# 创建测试文件，包含 package 声明
echo "package models" > internal/models/ngen_test.go
```

### 4. 编写生成逻辑

在测试文件中编写代码生成逻辑，配置生成选项并设置要处理的模型。

**注意**: 在 Go 中，使用测试文件生成源代码是一种常见的做法。

创建 `internal/models/ngen_test.go`：
```go
package models

import (
    "testing"
    
    "github.com/yyle88/gormcngen"
    "github.com/yyle88/osexistpath/osmustexist"
    "github.com/yyle88/runpath/runtestpath"
)

//go:generate go test -run ^TestGenerate$
func TestGenerate(t *testing.T) {
    // 获取目标文件的绝对路径（ngen.go）
    absPath := osmustexist.FILE(runtestpath.SrcPath(t))
    t.Log(absPath)
    
    // 配置生成选项
    options := gormcngen.NewOptions().
        WithColumnClassExportable(true).
        WithColumnsMethodRecvName("c").
        WithColumnsCheckFieldType(true)
    
    // 定义要生成的模型
    models := []interface{}{
		&User{},
	}
    
    // 创建配置并执行生成
    cfg := gormcngen.NewConfigs(models, options, absPath)
    cfg.Gen()
}
```

### 5. 执行生成

运行测试来触发代码生成，生成的代码将自动写入到目标文件中：

```bash
# 整理依赖
go mod tidy

# 运行生成测试
cd internal/models
go test -v ./...
```

🎉 **生成完成！** 您的 `ngen.go` 文件现在包含了生成的列访问方法。

生成的 `ngen.go` 将包含：

```go
// AUTO-GENERATED - DO NOT EDIT
// Generated by gormcngen

func (c *User) Columns() *UserColumns {
    return &UserColumns{
        ID:       gormcnm.Cnm(c.ID, "id"),
        Username: gormcnm.Cnm(c.Username, "username"),
        Email:    gormcnm.Cnm(c.Email, "email"),
        Age:      gormcnm.Cnm(c.Age, "age"),
        IsActive: gormcnm.Cnm(c.IsActive, "is_active"),
    }
}

type UserColumns struct {
    gormcnm.ColumnOperationClass
    ID       gormcnm.ColumnName[uint]
    Username gormcnm.ColumnName[string]
    Email    gormcnm.ColumnName[string]
    Age      gormcnm.ColumnName[int]
    IsActive gormcnm.ColumnName[bool]
}
```

🚀 **设置完成！** 您现在拥有了模型的类型安全列访问方法。

### 6. 在业务逻辑中使用

现在在编写业务代码时，您可以使用生成的类型安全列方法进行数据库查询：

```go
var user User
cls := user.Columns()

// 完美的类型安全，零样板代码
err := db.Where(cls.Username.Eq("alice")).
         Where(cls.Age.Gte(18)).
         Where(cls.IsActive.Eq(true)).
         First(&user).Error
```

✨ **这种方法的好处：**
- **编译时保护**: 列名拼写错误将变成编译错误
- **IDE 智能提示**: 完整的自动补全和重构支持
- **零样板代码**: 无需手动管理列名
- **始终同步**: 生成的代码与您的模型自动保持同步


### 高级用法

```go
// 基础配置（匹配内部示例）
options := gormcngen.NewOptions().
    WithColumnClassExportable(true).           // 生成导出的 ExampleColumns 结构体
    WithEmbedColumnOperations(false)           // 不嵌入操作方法

// 中文字段名支持
chineseOptions := gormcngen.NewOptions().
    WithUseTagName(true).                      // 使用 cnm 标签值作为字段名
    WithTagKeyName("cnm").                     // 指定 'cnm' 作为标签键
    WithColumnClassExportable(true)

// 高级功能（来自 example6）
advancedOptions := gormcngen.NewOptions().
    WithColumnClassExportable(true).           // 导出结构体名称
    WithColumnsMethodRecvName("one").          // 自定义接收器名称
    WithColumnsCheckFieldType(true).           // 类型检查（推荐）
    WithIsGenFuncTableColumns(true)            // 生成 TableColumns 函数

// 批量处理多个模型
allModels := []interface{}{&User{}, &Product{}, &Order{}, &Customer{}}
configs := gormcngen.NewConfigs(allModels, options, "models_gen.go")
configs.Gen()
```

## 高级功能

### 多语言字段支持

```go
type Product struct {
    ID          uint          `gorm:"primaryKey"`
    Name        string        `gorm:"size:255;not null" cnm:"V产品名称"`
    Price       decimal.Decimal `gorm:"type:decimal(10,2)"`
    CategoryID  uint          `gorm:"index"`
    CreatedAt   time.Time     `gorm:"autoCreateTime"`
    UpdatedAt   time.Time     `gorm:"autoUpdateTime"`
}
```

**生成结果：**

```go
type ProductColumns struct {
    gormcnm.ColumnOperationClass
    ID        gormcnm.ColumnName[uint]
    Name      gormcnm.ColumnName[string]      // 映射到 "name"
    V产品名称   gormcnm.ColumnName[string]      // 中文字段别名，映射到 Name 字段
    Price     gormcnm.ColumnName[decimal.Decimal]
    CategoryID gormcnm.ColumnName[uint]
    CreatedAt gormcnm.ColumnName[time.Time]
    UpdatedAt gormcnm.ColumnName[time.Time]
}

func (*Product) Columns() *ProductColumns {
    return &ProductColumns{
        ID:        "id",
        Name:      "name",
        V产品名称:   "name",           // 中文别名指向同一列
        Price:     "price",
        CategoryID: "category_id",
        CreatedAt: "created_at",
        UpdatedAt: "updated_at",
    }
}
```

**使用中文字段名进行查询：**

通过生成的中文字段别名，你可以使用母语编写查询语句：

```go
var product Product
var cls = product.Columns()

// 使用中文字段名查询 - 相同的数据库列，不同的 Go 字段名
if err := db.Where(cls.V产品名称.Eq("iPhone")).
    Where(cls.Price.Gte(5000.00)).
    First(&product).Error; err != nil {
    panic(errors.WithMessage(err, "未找到产品"))
}

fmt.Println("找到产品:", product.Name)
```

这让开发者可以用母语编写更易读的代码，同时保持完整的类型安全和数据库兼容性。

### Go Generate 集成

创建生成脚本：

**scripts/generate_columns.go:**
```go
package main

import (
    "github.com/yyle88/gormcngen"
    "your-project/models"
)

func main() {
    models := []interface{}{&models.User{}}
    options := gormcngen.NewOptions()
    configs := gormcngen.NewConfigs(models, options, "models/user_columns_gen.go")
    configs.Gen()
}
```

然后在模型文件中使用：

```go
//go:generate go run scripts/generate_columns.go

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex"`
    Email    string `gorm:"index"`
}
```

## 与 GORM 仓储模式集成

```go
// 生成的列与 gormrepo 无缝协作
repo := gormrepo.NewRepo(gormclass.Use(&Product{}))

products, total, err := repo.Repo(db).FindPageAndCount(
    func(db *gorm.DB, cls *ProductColumns) *gorm.DB {
        // 可以使用英文字段名
        return db.Where(cls.Name.Like("%computer%")).
               Where(cls.Price.Between(1000, 5000))
        // 或使用中文别名字段访问同一列
        // return db.Where(cls.V产品名称.Like("%电脑%")).
        //        Where(cls.Price.Between(1000, 5000))
    },
    func(cls *ProductColumns) gormcnm.OrderByBottle {
        return cls.Price.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 20, Offset: 0},
)
```

---

**通过使用 `gormcngen`，你可以轻松自动生成 `Columns()` 方法，进而用任何语言编写简单的查询语句。**

---

## 示例

查看 [examples](internal/examples) 和 [demos](internal/demos) 目录获取：
- 基础模型生成示例
- 中文字段处理示例
- 批量模型处理示例
- 自定义配置示例
- 真实数据库操作示例

## 相比手动列定义的优势

| 方面 | 手动定义 | GORMCNGEN |
|------|----------|-----------|
| **设置时间** | ⏰ 数小时手动输入 | ⚡ 编程 API 几秒钟 |
| **准确性** | ❌ 容易拼写错误 | ✅ 100% 准确的 AST 解析 |
| **同步性** | ❌ 需要手动更新 | ✅ 始终与模型同步 |
| **类型安全** | 🟡 依赖手动准确性 | ✅ 完美的类型保持 |
| **嵌入字段** | ❌ 复杂的手动处理 | ✅ 自动检测 |
| **原生语言** | ❌ 手动标签映射 | ✅ 智能标签处理 |
| **大型代码库** | 😫 维护噩梦 | 🚀 轻松扩展 |
| **团队生产力** | 🐌 缓慢且易错 | ⚡ 快速可靠 |

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormcngen.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcngen)

---

## 🔗 相关项目

- 🏗️ **[gormcnm](https://github.com/yyle88/gormcnm)** - 类型安全列基础包
- 🤖 **[gormcngen](https://github.com/yyle88/gormcngen)** - 智能代码生成（本包）
- 🏢 **[gormrepo](https://github.com/yyle88/gormrepo)** - 企业仓储模式
- 🌍 **[gormmom](https://github.com/yyle88/gormmom)** - 原生语言编程支持