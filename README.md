[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcngen/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcngen/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcngen)](https://pkg.go.dev/github.com/yyle88/gormcngen)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcngen/master.svg)](https://coveralls.io/github/yyle88/gormcngen?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcngen.svg)](https://github.com/yyle88/gormcngen/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcngen)](https://goreportcard.com/report/github.com/yyle88/gormcngen)

# gormcngen

`gormcngen`: Provides a Columns() Function to Retrieve Column Names for GORM Models

Like `MyBatis Plus` in the Java ecosystem, which allows developers to dynamically retrieve column names using expressions like `Example::getName`.

Like `SQLAlchemy` in the Python ecosystem, which allows developers to access column names using a `class` function, like `Example.name`.

`gormcngen` also brings **type-safe** column referencing to Go models.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Key Features

### 🔍 AST-Level Precision
- **Deep model analysis**: Parses struct fields, tags, and embedded types
- **GORM tag extraction**: Auto detects column names, types, and constraints
- **Embedded field support**: Handles `gorm.Model` and custom embedded structs
- **Type preservation**: Maintains exact Go types in generated code

### 🚀 Smart Code Generation
- **Perfect synchronization**: Generated code always matches your models
- **Custom column names**: Respects `gorm:"column:name"` tags
- **Native language support**: Works with `cnm:"中文名"` tags for international development
- **Incremental updates**: Regenerates just what changed

### 🛠️ Developer Experience
- **Simple programming API**: Easy-to-use Go API for immediate results
- **IDE integration**: Generated code provides full IntelliSense support
- **Build system friendly**: Easy integration with `go:generate` directives
- **Version control safe**: Deterministic output for clean diffs

### 🏢 Enterprise Ready
- **Large codebase support**: Handles hundreds of models efficiently
- **Custom naming conventions**: Configurable output patterns
- **Validation and safety**: Built-in checks prevent invalid generation
- **Documentation generation**: Auto-generated comments explain column mappings

## 🏗️ Ecosystem Position

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

**gormcngen** serves as the **code generation engine** that bridges your models with the type-safe foundation layer.

## Install

```bash
go get github.com/yyle88/gormcngen
```

## Quick Start

### 1. Define Your GORM Model

```go
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex;size:100"`
    Email    string `gorm:"index;size:255"`
    Age      int    `gorm:"check:age >= 0"`
    IsActive bool   `gorm:"default:true"`
}
```

### 2. Generate Type-Safe Columns

```go
package main

import (
    "github.com/yyle88/gormcngen"
)

func main() {
    // Configure generation options
    models := []interface{}{&User{}}
    options := gormcngen.NewOptions()
    outputPath := "user_columns_gen.go"
    
    // Generate code
    configs := gormcngen.NewConfigs(models, options, outputPath)
    configs.Gen()
}
```

### 3. Generated Code (Automatically!)

```go
// AUTO-GENERATED - DO NOT EDIT
// Generated by gormcngen

type UserColumns struct {
    gormcnm.ColumnOperationClass
    // The column names and types of the model's columns
    ID       gormcnm.ColumnName[uint]
    Username gormcnm.ColumnName[string]
    Email    gormcnm.ColumnName[string]
    Age      gormcnm.ColumnName[int]
    IsActive gormcnm.ColumnName[bool]
}

func (*User) Columns() *UserColumns {
    return &UserColumns{
        ID:       "id",
        Username: "username",
        Email:    "email",
        Age:      "age",
        IsActive: "is_active",
    }
}
```

### 4. Use Type-Safe Queries

```go
var user User
cls := user.Columns()

// Perfect type safety with zero boilerplate
err := db.Where(cls.Username.Eq("alice")).
         Where(cls.Age.Gte(18)).
         Where(cls.IsActive.Eq(true)).
         First(&user).Error
```

### Advanced Usage

```go
// Basic configuration (matches internal examples)
options := gormcngen.NewOptions().
    WithColumnClassExportable(true).           // Generate exported ExampleColumns struct
    WithEmbedColumnOperations(false)           // Don't embed operation methods

// Chinese field name support
chineseOptions := gormcngen.NewOptions().
    WithUseTagName(true).                      // Use cnm tag values as field names  
    WithTagKeyName("cnm").                     // Specify 'cnm' as the tag key
    WithColumnClassExportable(true)

// Advanced features (from example6)
advancedOptions := gormcngen.NewOptions().
    WithColumnClassExportable(true).           // Exported struct names
    WithColumnsMethodRecvName("one").          // Custom receiver name
    WithColumnsCheckFieldType(true).           // Type checking (recommended)
    WithIsGenFuncTableColumns(true)            // Generate TableColumns function

// Batch processing multiple models
allModels := []interface{}{&User{}, &Product{}, &Order{}, &Customer{}}
configs := gormcngen.NewConfigs(allModels, options, "models_gen.go")
configs.Gen()
```

## Advanced Features

### Multi-Language Field Support

The `cnm` tag allows you to define Chinese aliases for field names, which are generated as additional struct fields:

```go
type Product struct {
    ID          uint          `gorm:"primaryKey"`
    Name        string        `gorm:"size:255;not null" cnm:"产品名称"`
    Price       decimal.Decimal `gorm:"type:decimal(10,2)"`
    CategoryID  uint          `gorm:"index"`
    CreatedAt   time.Time     `gorm:"autoCreateTime"`
    UpdatedAt   time.Time     `gorm:"autoUpdateTime"`
}
```

**Generated Result:**

```go
type ProductColumns struct {
    gormcnm.ColumnOperationClass
    // The column names and types of the model's columns
    ID       gormcnm.ColumnName[uint]
    Name     gormcnm.ColumnName[string]           // Maps to "name"
    V产品名称   gormcnm.ColumnName[string]           // Chinese alias for Name field  
    Price    gormcnm.ColumnName[decimal.Decimal]
    CategoryID gormcnm.ColumnName[uint]
    CreatedAt gormcnm.ColumnName[time.Time]
    UpdatedAt gormcnm.ColumnName[time.Time]
}

func (*Product) Columns() *ProductColumns {
    return &ProductColumns{
        ID:       "id",
        Name:     "name",
        V产品名称:   "name",      // Chinese alias for same column
        Price:    "price",
        CategoryID: "category_id",
        CreatedAt: "created_at",
        UpdatedAt: "updated_at",
    }
}
```

**Using Chinese Field Names in Queries:**

With the generated Chinese aliases, you can write queries using your native language:

```go
var product Product
var cls = product.Columns()

// Query using Chinese field names - same database column, different Go field name
if err := db.Where(cls.V产品名称.Eq("iPhone")).
    Where(cls.Price.Gte(5000.00)).
    First(&product).Error; err != nil {
    panic(errors.WithMessage(err, "product not found"))
}

fmt.Println("Found product:", product.Name)
```

This allows developers to write more readable code in their native language while maintaining full type safety and database compatibility.

### Go Generate Integration

Create a generation script:

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

Then use in your model files:

```go
//go:generate go run scripts/generate_columns.go

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex"`
    Email    string `gorm:"index"`
}
```

## Integration with GORM Repository Pattern

```go
// Generated columns work seamlessly with gormrepo
repo := gormrepo.NewRepo(gormclass.Use(&Product{}))

products, total, err := repo.Repo(db).FindPageAndCount(
    func(db *gorm.DB, cls *ProductColumns) *gorm.DB {
        // Can use English field name
        return db.Where(cls.Name.Like("%computer%")).
               Where(cls.Price.Between(1000, 5000))
        // Or use Chinese alias field for same column
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

**This is a straightforward way to install and use `gormcngen` to generate the `Columns()` method for GORM models, allowing you to easily build queries with column names in any language.**

---

## Examples

See [examples](internal/examples) and [demos](internal/demos) directories for:
- Basic model generation examples
- Chinese field handling examples
- Batch model processing examples
- Custom configuration examples
- Real database operation examples

## Benefits Over Manual Column Definitions

| Aspect | Manual Definitions | GORMCNGEN |
|--------|-------------------|-----------|
| **Setup Time** | ⏰ Hours of manual typing | ⚡ Seconds with programming API |
| **Accuracy** | ❌ Prone to typos and mistakes | ✅ 100% accurate AST parsing |
| **Synchronization** | ❌ Manual updates required | ✅ Always in sync with models |
| **Type Safety** | 🟡 Depends on manual accuracy | ✅ Perfect type preservation |
| **Embedded Fields** | ❌ Complex manual handling | ✅ Automatic detection |
| **Native Language** | ❌ Manual tag mapping | ✅ Intelligent tag processing |
| **Large Codebases** | 😫 Maintenance nightmare | 🚀 Scales effortlessly |
| **Team Productivity** | 🐌 Slow and error-prone | ⚡ Fast and reliable |

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share your use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize by reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo for new releases and features
- 🌟 **Success stories?** Share how this package improved your workflow
- 💬 **General feedback?** All suggestions and comments are welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage interface).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement your changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation for user-facing changes and use meaningful commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormcngen.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcngen)

---

## 🔗 Related Projects

- 🏗️ **[gormcnm](https://github.com/yyle88/gormcnm)** - Type-safe column foundation
- 🤖 **[gormcngen](https://github.com/yyle88/gormcngen)** - Smart code generation (this package)
- 🏢 **[gormrepo](https://github.com/yyle88/gormrepo)** - Enterprise repository pattern
- 🌍 **[gormmom](https://github.com/yyle88/gormmom)** - Native language programming