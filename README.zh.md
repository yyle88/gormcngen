# `gormcngen`: 赋予 GORM 模型使用 Columns() 获取列名的函数

就像 Java 生态系统中的 `MyBatis Plus`，它允许开发人员使用像 `Example::getName` 这样的表达式获取列名。

就像 Python 生态系统中的 `SQLAlchemy`，它允许开发人员使用像 `Example.name` 这样的 class 函数获得列名。

`gormcngen` 也为 Go 模型赋予 **类型安全** 的列引用功能。

## 英文文档

[ENGLISH README](README.md)

## 安装

```bash
go get github.com/yyle88/gormcngen
```

## 示例使用

### 1. 首先定义模型

假设你有如下模型：

```go
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}
```

### 2. 自动生成 `Columns()` 方法

使用 `gormcngen`，它会自动为你的模型生成 `Columns()` 方法：

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

### 3. 使用生成的 `Columns()` 进行查询

你现在可以轻松地使用生成的 `Columns()` 方法来构建查询：

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

### 4. 自定义列名示例（如使用中文）

如果你的模型包含自定义的列名（例如使用中文），它的使用方法是一样的：

```go
type Demo struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);" cnm:"V名称"`
	Type string `gorm:"type:varchar(100);" cnm:"V类型"`
}
```

生成的代码：

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

这样，你就可以在查询时使用母语（如中文）：

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

通过使用 `gormcngen`，你可以轻松自动生成 `Columns()` 方法，进而用任何语言编写简单的查询语句。

---

## 完整样例

[完整样例](internal/demos)

## 设计思路

[旧版说明](internal/docs/README_OLD_DOC.zh.md)

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
