[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcngen/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcngen/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcngen)](https://pkg.go.dev/github.com/yyle88/gormcngen)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcngen/master.svg)](https://coveralls.io/github/yyle88/gormcngen?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcngen.svg)](https://github.com/yyle88/gormcngen/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcngen)](https://goreportcard.com/report/github.com/yyle88/gormcngen)

# `gormcngen`: Provides a Columns() Function to Retrieve Column Names for GORM Models

Like `MyBatis Plus` in the Java ecosystem, which allows developers to dynamically retrieve column names using expressions like `Example::getName`.

Like `SQLAlchemy` in the Python ecosystem, which allows developers to access column names using a `class` function, like `Example.name`.

`gormcngen` also brings **type-safe** column referencing to Go models.

## CHINESE README

[中文说明](README.zh.md)

## Installation

```bash
go get github.com/yyle88/gormcngen
```

## Example Usage

### 1. Define Your Model

For example, let's say you have the following model:

```go
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}
```

### 2. Automatically Generate the `Columns()` Method

Using `gormcngen`, it will automatically generate the `Columns()` method for your model:

```go
func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		Name: "name",
		Type: "type",
		Rank: "rank",
	}
}

type ExampleColumns struct {
	Name gormcnm.ColumnName[string]
	Type gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}
```

### 3. Querying with the Generated `Columns()`

Now you can easily use the generated `Columns()` method to build queries:

```go
var res Example
var cls = res.Columns()

if err := db.Where(cls.Name.Eq("abc")).
    Where(cls.Type.Eq("xyz")).
    Where(cls.Rank.Gt(100)).
    Where(cls.Rank.Lt(200)).
    First(&res).Error; err != nil {
    panic(errors.WithMessage(err, "wrong"))
}

fmt.Println(res)
```

### 4. Example with Custom Column Names

If your model contains custom column names (like using Chinese), it works similarly:

```go
type Demo struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);" cnm:"V名称"`
	Type string `gorm:"type:varchar(100);" cnm:"V类型"`
}
```

Generated code:

```go
func (*Demo) Columns() *DemoColumns {
	return &DemoColumns{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		V名称:      "name",
		V类型:      "type",
	}
}

type DemoColumns struct {
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	V名称      gormcnm.ColumnName[string]
	V类型      gormcnm.ColumnName[string]
}
```

With this, you can use your native language for column names when querying:

```go
var demo Demo
var cls = demo.Columns()

if err := db.Where(cls.V名称.Eq("测试")).
    Where(cls.V类型.Eq("类型A")).
    First(&demo).Error; err != nil {
    panic(errors.WithMessage(err, "wrong"))
}

fmt.Println(demo)
```

---

This is a more straightforward explanation of how to install and use `gormcngen` to generate the `Columns()` method for GORM models, allowing you to easily build queries with column names in any language.

---

## Demos

[demos](internal/demos)

## Design Ideas

[README OLD DOC](internal/docs/README_OLD_DOC.en.md)

---

## License

`gormcngen` is open-source and released under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `gormcngen`!** 🎉

Give me stars. Thank you!!!

## See stars
[![see stars](https://starchart.cc/yyle88/gormcngen.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcngen)
