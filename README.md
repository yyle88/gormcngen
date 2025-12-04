[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcngen/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcngen/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcngen)](https://pkg.go.dev/github.com/yyle88/gormcngen)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcngen/main.svg)](https://coveralls.io/github/yyle88/gormcngen?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcngen.svg)](https://github.com/yyle88/gormcngen/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcngen)](https://goreportcard.com/report/github.com/yyle88/gormcngen)

# gormcngen

`gormcngen`: Provides a Columns() Function to Retrieve Column Names from GORM Models

Like `MyBatis Plus` in the Java ecosystem, which allows developers to get column names using expressions like `Example::getName`.

Like `SQLAlchemy` in the Python ecosystem, which allows developers to access column names using a `class` function, like `Example.name`.

`gormcngen` also brings **type-safe** column referencing to Go models.

---

## Ecosystem

![GORM Type-Safe Ecosystem](https://github.com/yyle88/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

---

## Core Features

### 🔍 AST-Grade Precision
- **Deep model analysis**: Parses struct fields, tags, and embedded types
- **GORM tag extraction**: Auto detects column names, types, and constraints
- **Embedded field support**: Handles `gorm.Model` and custom embedded structs
- **Type preservation**: Maintains exact Go types in generated code

### 🚀 Smart Code Generation
- **Perfect synchronization**: Generated code always matches the models
- **Custom column names**: Respects `gorm:"column:name"` tags
- **Native language support**: Works with `cnm:"中文名"` tags to enable multi-language development
- **Smart updates**: Regenerates just what changed

### 🛠️ Development Experience
- **Simple programming API**: Simple Go API with immediate results
- **IDE integration**: Generated code provides complete IntelliSense support
- **Build system compatible**: Simple integration with `go:generate` directives
- **Version management safe**: Deterministic output with clean diffs

### 🏢 Enterprise Grade
- **Large codebase support**: Handles hundreds of models with ease
- **Custom naming conventions**: Configurable output patterns
- **Validation and protection**: Built-in checks prevent invalid generation
- **Documentation generation**: Auto-generated comments explain column mappings

## Installation

```bash
go get github.com/yyle88/gormcngen
```

## Important Note

**gormcngen is a Go package, not a CLI application.** It requires a test-code-driven generation workflow.

## Quick Start

### 1. Create Project Structure

Set up the basic project structure and create dedicated DIR to hold models and generated code:

```bash
# Create models DIR
mkdir -p internal/models
```

### 2. Define GORM Models

Define data models - gormcngen generates column access methods from these models:

Create `internal/models/models.go`:
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

### 3. Create Generation Files

Create the target file to hold generated code and the test file containing generation logic:

```bash
# Create target file for generated code with package declaration
echo "package models" > internal/models/ngen.go

# Create test file for generation logic with package declaration
echo "package models" > internal/models/ngen_test.go
```

### 4. Write Generation Logic

Write the code generation logic in the test file, configure generation options and set the models to process.

**Note**: In Go, using test files to generate source code is a common practice.

Create `internal/models/ngen_test.go`:
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
    // Get absolute path to target file (ngen.go)
    absPath := osmustexist.FILE(runtestpath.SrcPath(t))
    t.Log(absPath)
    
    // Configure generation options
    options := gormcngen.NewOptions().
        WithColumnClassExportable(true).
        WithColumnsMethodRecvName("c").
        WithColumnsCheckFieldType(true)
    
    // Define models to generate for
    models := []interface{}{
		&User{},
	}
    
    // Create config and generate
    cfg := gormcngen.NewConfigs(models, options, absPath)
    cfg.WithIsGenPreventEdit(true)  // Add "DO NOT EDIT" warning headers (default: true)
    cfg.WithGeneratedFromPos(gormcngen.GetGenPosFuncMark(0))  // Show generation source location (default: show)
    cfg.Gen()
}
```

### 5. Execute Generation

Run the test to initiate code generation - the generated code gets auto written to the target file:

```bash
# Clean up dependencies
go mod tidy

# Run generation test
cd internal/models
go test -v ./...
```

🎉 **Generation Complete!** The `ngen.go` file now contains the generated column access methods.

The generated `ngen.go` contains:

```go
// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: ngen_test.go:20 -> models.TestGenerate
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

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

🚀 **Setup Complete!** You now have type-safe column access methods to work with models.

### 6. Use in Business Logic

Now when writing business code, you can use the generated type-safe column methods to build database queries:

```go
var user User
cls := user.Columns()

// Perfect type protection with zero boilerplate
err := db.Where(cls.Username.Eq("alice")).
         Where(cls.Age.Gte(18)).
         Where(cls.IsActive.Eq(true)).
         First(&user).Error
```

✨ **Benefits of this approach:**
- **Compile-time protection**: Typos in column names become compilation errors
- **IDE intelligence**: Complete autocomplete and refactoring support  
- **Zero boilerplate**: No hand-written column name management required
- **Always synchronized**: Generated code stays in sync with the models auto

### Advanced Usage

```go
// Basic configuration (matches internal examples)
options := gormcngen.NewOptions().
    WithColumnClassExportable(true).           // Generate exported ExampleColumns struct
    WithEmbedColumnOperations(false)           // Don't embed operation methods

// Chinese field name support
chineseOptions := gormcngen.NewOptions().
    WithUseTagName(true).                      // Use cnm tag values as field names  
    WithTagKeyName("cnm").                     // Set 'cnm' as the tag name
    WithColumnClassExportable(true)

// Advanced features (from example6)
advancedOptions := gormcngen.NewOptions().
    WithColumnClassExportable(true).           // Exported struct names
    WithColumnsMethodRecvName("one").          // Custom method argument
    WithColumnsCheckFieldType(true).           // Type checking (recommended)
    WithIsGenFuncTableColumns(true)            // Generate TableColumns function

// Batch processing multiple models
allModels := []interface{}{&User{}, &Product{}, &Item{}, &Client{}}
configs := gormcngen.NewConfigs(allModels, options, "models_gen.go")
configs.WithIsGenPreventEdit(true)  // Add "DO NOT EDIT" headers to generated files
configs.WithGeneratedFromPos(gormcngen.GetGenPosFuncMark(0))  // Show generation source location (default: show)
configs.Gen()
```

## Advanced Features

### Multi-Language Field Support

The `cnm` tag lets you define Chinese aliases to use as field names, which are generated as extra struct fields:

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

With the generated Chinese aliases, you can write queries using native language:

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

This allows developers to write more readable code in native language while maintaining complete type protection and database support.

### Go Generate Integration

Create a generation script:

**scripts/generate_columns.go:**
```go
package main

import (
    "github.com/yyle88/gormcngen"
    "project-name/models"
)

func main() {
    models := []interface{}{&models.User{}}
    options := gormcngen.NewOptions()
    configs := gormcngen.NewConfigs(models, options, "models/user_columns_gen.go")
    configs.Gen()
}
```

Then use in the target files:

```go
//go:generate go run scripts/generate_columns.go

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex"`
    Email    string `gorm:"index"`
}
```

## Integration with GORM Repo Pattern

```go
// Generated columns integrate with gormrepo
repo := gormrepo.NewRepo(gormclass.Use(&Product{}))

products, total, err := repo.Repo(db).FindPageAndCount(
    func(db *gorm.DB, cls *ProductColumns) *gorm.DB {
        // Can use English field name
        return db.Where(cls.Name.Like("%computer%")).
               Where(cls.Price.Between(1000, 5000))
        // Also supports Chinese alias field access to same column
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

**This is a simple method to setup and use `gormcngen` to generate the `Columns()` function with GORM models, enabling you to build queries with column names in different languages.**

---

## Examples

See [examples](internal/examples) and [demos](internal/demos) directories for:
- Basic generation examples
- Chinese field handling examples
- Batch processing examples
- Custom configuration examples
- Database operation examples

## Benefits Versus Static Column Definitions

| Aspect | Static Definitions | GORMCNGEN |
|--------|-------------------|-----------|
| **Setup Time** | ⏰ Hours of hand-written typing | ⚡ Seconds with programming API |
| **Precision** | ❌ Prone to typos and mistakes | ✅ 100% accurate AST parsing |
| **Synchronization** | ❌ Hand-written updates required | ✅ Always in sync with models |
| **Type Protection** | 🟡 Depends on hand-written precision | ✅ Perfect type preservation |
| **Embedded Fields** | ❌ Complex hand-written handling | ✅ Automatic detection |
| **Native Language** | ❌ Hand-written tag mapping | ✅ Intelligent tag processing |
| **Large Codebases** | 😫 Maintenance nightmare | 🚀 Works at scale |
| **Team Performance** | 🐌 Slow and error-prone | ⚡ Fast and reliable |

---

## Related Projects

Explore the complete GORM ecosystem with these integrated packages:

### Core Ecosystem

- **[gormcnm](https://github.com/yyle88/gormcnm)** - GORM foundation providing type-safe column operations and query builders
- **[gormcngen](https://github.com/yyle88/gormcngen)** - Code generation tool using AST for type-safe GORM operations (this project)
- **[gormrepo](https://github.com/yyle88/gormrepo)** - Repository pattern implementation with GORM best practices
- **[gormmom](https://github.com/yyle88/gormmom)** - Native language GORM tag generation engine with smart column naming
- **[gormzhcn](https://github.com/go-zwbc/gormzhcn)** - Complete Chinese programming interface with GORM

Each package targets different aspects of GORM development, from localization to type safety and code generation.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormcngen.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcngen)

