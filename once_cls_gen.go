package gormcngen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/tern"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
	"gorm.io/gorm/schema"
)

// SchemaConfig Configuration of generating column methods and structures.
// SchemaConfig 根据模型生成列方法和结构的配置。
type SchemaConfig struct {
	sch        *schema.Schema // Parsed schema from the model.// 结构体模型对应的数据表结构。
	structName string         // Name of the generated structure.// 生成的结构体名称。
	methodName string         // Name of the generated method.// 生成的方法名称。
	options    *Options       // Additional configuration options.// 额外的配置选项。
}

// NewSchemaConfig Creates a Config instance for the given destination model and options.
// NewSchemaConfig 为指定的目标模型和选项创建 Config 实例。
func NewSchemaConfig(object interface{}, options *Options) *SchemaConfig {
	sch := done.VCE(schema.Parse(object, &sync.Map{}, &schema.NamingStrategy{
		SingularTable: false,
		NoLowerCase:   false,
	})).Nice()

	ShowSchemaEnglish(sch)
	ShowSchemaChinese(sch)

	const structNameSuffix = "Columns"
	const methodName = "Columns"

	var structName string
	if !options.columnClassExportable {
		structName = utils.ConvertToUnexportable(sch.Name) + structNameSuffix
	} else {
		structName = sch.Name + structNameSuffix // 通常定义的结构体名称是导出的
	}

	return NewConfig(sch, structName, methodName, options)
}

// Config Configuration of generating column methods and structures.
// Config 根据模型生成列方法和结构的配置。
type Config = SchemaConfig

// NewConfig Creates a new Config instance with the provided schema, struct name, method name, and options.
// NewConfig 创建一个新的 Config 实例，使用提供的 schema、结构体名称、方法名称和选项。
func NewConfig(sch *schema.Schema, structName string, methodName string, options *Options) *Config {
	return &Config{
		sch:        sch,
		structName: structName,
		methodName: methodName,
		options:    options,
	}
}

// ColumnsMethodStructOutput Structure representing the generated method and struct code with package imports.
// ColumnsMethodStructOutput 表示生成的方法和结构体代码，以及涉及的包导入信息。
type ColumnsMethodStructOutput struct {
	methodCode string          // Code for the generated method.// 生成的方法代码。
	structCode string          // Code for the generated structure.// 生成的结构体代码。
	pkgImports map[string]bool // Package imports required by the generated code.// 生成代码需要的包导入。
}

func (x *ColumnsMethodStructOutput) GetMethodCode() string {
	return x.methodCode
}

func (x *ColumnsMethodStructOutput) GetStructCode() string {
	return x.structCode
}

func (x *ColumnsMethodStructOutput) GetPkgImports() map[string]bool {
	return x.pkgImports
}

// Generate Generates the column method and struct based on the configuration.
// Generate 根据配置生成列方法和结构。
func (c *Config) Generate() *ColumnsMethodStructOutput {
	return c.Gen()
}

// Gen Generates the column method and struct based on the configuration.
// Gen 根据配置生成列方法和结构。
func (c *Config) Gen() *ColumnsMethodStructOutput {
	structPtx := utils.NewPTX()
	structPtx.Println(fmt.Sprintf("type %s struct{", c.structName))

	methodPtx := utils.NewPTX()
	methodPtx.Println(fmt.Sprintf("func (%s*%s) %s() *%s {", c.options.columnsMethodRecvName, c.sch.Name, c.methodName, c.structName))

	operationClass := reflect.TypeOf(gormcnm.ColumnOperationClass{})
	pkgNameGormCnm := filepath.Base(operationClass.PkgPath())

	var pkgImports = map[string]bool{
		operationClass.PkgPath(): true,
	}

	const indentPrefix = "   " // 用于代码对齐的缩进（3个空格）

	if c.options.embedColumnOperations {
		structPtx.Println(indentPrefix, "// Embedding operation functions make it easy to use // 继承操作函数便于使用")
		structPtx.Println(indentPrefix, fmt.Sprintf("%s.%s", pkgNameGormCnm, operationClass.Name()))
	}
	structPtx.Println(indentPrefix, "// The column names and types of the model's columns // 模型各列的列名和类型")

	methodPtx.Println(fmt.Sprintf("	return &%s{", c.structName))
	for _, field := range c.sch.Fields {
		var columnGoTypeName string
		if pkgPath := field.FieldType.PkgPath(); pkgPath == c.sch.ModelType.PkgPath() { // 如果在同一个包里，仅使用类型名
			columnGoTypeName = field.FieldType.Name()
		} else {
			if pkgPath != "" {
				pkgImports[pkgPath] = true
			}
			columnGoTypeName = field.FieldType.String() // 使用完整类型名
		}
		newStructFieldName, ok := c.resolveNewFieldName(field)
		if !ok {
			continue // 某些场景下不需要获得列名
		}

		structPtx.Println(indentPrefix, newStructFieldName, fmt.Sprintf("%s.ColumnName[%s]", pkgNameGormCnm, columnGoTypeName))

		dbColumnName := tern.BFF(c.options.columnsMethodRecvName != "" && c.options.columnsCheckFieldType, func() string {
			return fmt.Sprintf(`%s.Cnm(%s.%s, "%s")`, pkgNameGormCnm, c.options.columnsMethodRecvName, field.Name, field.DBName)
		}, func() string {
			return `"` + field.DBName + `"`
		})

		methodPtx.Println(indentPrefix, indentPrefix, fmt.Sprintf("%s:%s,", newStructFieldName, dbColumnName))
	}
	structPtx.Println("}")
	methodPtx.Println("	}")
	methodPtx.Println("}")

	structCode := strings.TrimSpace(structPtx.String())
	methodCode := strings.TrimSpace(methodPtx.String())

	zaplog.SUG.Debug("---", "\n", methodCode)
	zaplog.SUG.Debug("---", "\n", structCode)
	zaplog.SUG.Debug("---", "\n", neatjsons.S(pkgImports))

	return &ColumnsMethodStructOutput{
		methodCode: methodCode,
		structCode: structCode,
		pkgImports: pkgImports,
	}
}

// Resolves the new field name based on tags and options.
// 根据标签和选项解析新字段名称。
func (c *Config) resolveNewFieldName(field *schema.Field) (string, bool) {
	if c.options.useTagName {
		var tagKeyName = zerotern.VV(c.options.tagKeyName, "cnm")

		name, ok := field.Tag.Lookup(tagKeyName)
		if ok {
			if !utils.IsExportable(name) { // 确保字段名是导出的
				panic(erero.Errorf("name=%v is not exportable", name))
			}
			return name, true
		} else {
			if c.options.excludeUntaggedFields {
				return "", false
			}
			return field.Name, true
		}
	}
	return field.Name, true
}

// ShowSchemaEnglish Displays schema information including struct name, table name, and fields.
// ShowSchemaEnglish 显示模式结构信息，包括结构体名称、表名和字段信息。
func ShowSchemaEnglish(sch *schema.Schema) {
	fmt.Println("---")
	fmt.Println("schema_message", "Struct name:", sch.Name, "Table name:", sch.Table, "Fields: {")
	for _, field := range sch.Fields {
		fmt.Println("   ",
			"Go Field Name:", field.Name,
			" | ",
			"Go Type:", field.FieldType,
			" | ",
			"DB Field Name:", field.DBName,
			" | ",
			"DB Type:", field.DataType,
			" | ",
			"Go Tag:", field.Tag,
		)
	}
	fmt.Println("}")
	fmt.Println("---")
}

// ShowSchemaChinese Displays schema information including struct name, table name, and fields.
// ShowSchemaChinese 显示模式结构信息，包括结构体名称、表名和字段信息。
func ShowSchemaChinese(sch *schema.Schema) {
	fmt.Println("---")
	fmt.Println("schema_message", "结构体名称:", sch.Name, "表名:", sch.Table, "字段信息: {")
	for _, field := range sch.Fields {
		fmt.Println("   ",
			"Go字段名:", field.Name, // Go结构体成员名称
			" | ",
			"Go类型:", field.FieldType, // Go的数据类型
			" | ",
			"DB字段名:", field.DBName, // 数据表中的列名
			" | ",
			"DB类型:", field.DataType, //数据库中的数据类型
			" | ",
			"Go标签:", field.Tag,
		)
	}
	fmt.Println("}")
	fmt.Println("---")
}
