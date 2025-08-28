// Package gormcngen: Schema-based configuration and single-use code generation
// Handles individual model schema analysis and targeted code generation
// Provides precise control over column struct generation and method creation
//
// gormcngen: 基于 schema 的配置和单次使用代码生成
// 处理单个模型 schema 分析和针对性代码生成
// 提供对列结构体生成和方法创建的精确控制
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
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/tern"
	"github.com/yyle88/tern/zerotern"
	"github.com/yyle88/zaplog"
	"gorm.io/gorm/schema"
)

// SchemaConfig holds configuration for generating column methods and structures from GORM models
// Contains parsed schema information along with generation options and naming preferences
// Manages the complete lifecycle from schema analysis to code output generation
//
// SchemaConfig 保存从 GORM 模型生成列方法和结构体的配置
// 包含已解析的 schema 信息以及生成选项和命名偏好
// 管理从 schema 分析到代码输出生成的完整生命周期
type SchemaConfig struct {
	sch                    *schema.Schema // Parsed GORM schema from model struct // 从模型结构体解析的 GORM schema
	structName             string         // Generated column struct name // 生成的列结构体名称
	methodName             string         // Generated Columns() method name // 生成的 Columns() 方法名称
	methodNameTableColumns string
	options                *Options // Generation behavior configuration options // 生成行为配置选项
}

// NewSchemaConfig creates a SchemaConfig instance from a GORM model and generation options
// Parses the model structure using GORM schema parsing and applies naming strategies
// Initializes schema analysis, shows debug information, and configures generation parameters
// Returns a fully configured SchemaConfig ready for code generation
//
// NewSchemaConfig 从 GORM 模型和生成选项创建 SchemaConfig 实例
// 使用 GORM schema 解析器解析模型结构并应用命名策略
// 初始化 schema 分析，显示调试信息，并配置生成参数
// 返回一个完全配置的 SchemaConfig，准备进行代码生成
func NewSchemaConfig(object interface{}, options *Options) *SchemaConfig {
	sch := done.VCE(schema.Parse(object, &sync.Map{}, &schema.NamingStrategy{
		SingularTable: false, //这是gorm默认的
		NoLowerCase:   false, //这是gorm默认的
	})).Nice()

	ShowSchemaEnglish(sch)
	ShowSchemaChinese(sch)

	namingConfig := NewNamingConfig(sch, options)

	return NewConfig(sch, namingConfig, options)
}

type NamingConfig struct {
	StructName             string
	MethodNameColumns      string
	MethodNameTableColumns string
}

func NewNamingConfig(sch *schema.Schema, options *Options) *NamingConfig {
	const structNameSuffix = "Columns"
	structName := tern.BFF(options.columnClassExportable, func() string {
		return sch.Name + structNameSuffix // 通常定义的结构体名称是导出的
	}, func() string {
		return utils.ConvertToUnexportable(sch.Name) + structNameSuffix
	})

	namingConfig := &NamingConfig{
		StructName:             structName,
		MethodNameColumns:      "Columns",
		MethodNameTableColumns: "TableColumns",
	}
	return namingConfig
}

// Config Configuration of generating column methods and structures.
// Config 根据模型生成列方法和结构的配置。
type Config = SchemaConfig

// NewConfig Creates a new Config instance with the provided schema, struct name, method name, and options.
// NewConfig 创建一个新的 Config 实例，使用提供的 schema、结构体名称、方法名称和选项。
func NewConfig(sch *schema.Schema, namingConfig *NamingConfig, options *Options) *Config {
	return &Config{
		sch:                    sch,
		structName:             namingConfig.StructName,
		methodName:             namingConfig.MethodNameColumns,
		methodNameTableColumns: namingConfig.MethodNameTableColumns,
		options:                options,
	}
}

// GenOutput Structure representing the generated method and struct code with package imports.
// GenOutput 表示生成的方法和结构体代码，以及涉及的包导入信息。
type GenOutput struct {
	methodCode             string // Code for the generated method. // 生成的方法代码。
	methodTableColumnsCode string
	structCode             string          // Code for the generated structure. // 生成的结构体代码。
	pkgImports             map[string]bool // Package imports required. // 生成的代码需要导入的的包。
}

func (x *GenOutput) GetMethodCode() string {
	return x.methodCode
}

func (x *GenOutput) GetMethodTableColumnsCode() string {
	return x.methodTableColumnsCode
}

func (x *GenOutput) GetStructCode() string {
	return x.structCode
}

func (x *GenOutput) GetPkgImports() map[string]bool {
	return x.pkgImports
}

// Generate Generates the column method and struct based on the configuration.
// Generate 根据配置生成列方法和结构。
func (c *Config) Generate() *GenOutput {
	return c.Gen()
}

// Gen Generates the column method and struct based on the configuration.
// Gen 根据配置生成列方法和结构。
func (c *Config) Gen() *GenOutput {
	structPtx := utils.NewPTX()
	structPtx.Println(fmt.Sprintf("type %s struct{", c.structName))

	methodPtx := utils.NewPTX()
	methodPtx.Println(fmt.Sprintf("func (%s*%s) %s() *%s {", c.options.columnsMethodRecvName, c.sch.Name, c.methodName, c.structName))

	methodTableColumnsPtx := utils.NewPTX()
	methodTableColumnsPtx.Println(fmt.Sprintf("func (%s*%s) %s(decoration gormcnm.ColumnNameDecoration) *%s {", c.options.columnsMethodRecvName, c.sch.Name, c.methodNameTableColumns, c.structName))

	operationClass := reflect.TypeOf(gormcnm.ColumnOperationClass{})
	pkgNameGormCnm := filepath.Base(operationClass.PkgPath())

	var pkgImports = map[string]bool{
		operationClass.PkgPath(): true,
	}

	if c.options.embedColumnOperations {
		structPtx.Println("\t", "// Embedding operation functions make it easy to use // 继承操作函数便于使用")
		structPtx.Println("\t", fmt.Sprintf("%s.%s", pkgNameGormCnm, operationClass.Name()))
	}
	structPtx.Println("\t", "// The column names and types of the model's columns // 模型各列的列名和类型")

	if c.options.isGenFuncTableColumns {
		must.Nice(c.options.columnsMethodRecvName)
		must.True(c.options.columnsCheckFieldType)
		methodPtx.Println(fmt.Sprintf("	return %s.%s(gormcnm.NewPlainDecoration())", c.options.columnsMethodRecvName, c.methodNameTableColumns))
		methodTableColumnsPtx.Println(fmt.Sprintf("	return &%s{", c.structName))
	} else {
		methodPtx.Println(fmt.Sprintf("	return &%s{", c.structName))
		methodTableColumnsPtx.Println(fmt.Sprintf("	panic(\"METHOD %s.%s IS NOT IMPLEMENTED\")", c.structName, c.methodNameTableColumns))
	}
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
		structFieldName, ok := c.resolveNewFieldName(field)
		if !ok {
			continue // 某些场景下不需要获得列名
		}
		structPtx.Println("\t", structFieldName, fmt.Sprintf("%s.ColumnName[%s]", pkgNameGormCnm, columnGoTypeName))

		if c.options.isGenFuncTableColumns {
			must.Nice(c.options.columnsMethodRecvName)
			must.True(c.options.columnsCheckFieldType)
			valueColumnName := fmt.Sprintf(`%s.Cmn(%s.%s, "%s", decoration)`, pkgNameGormCnm, c.options.columnsMethodRecvName, field.Name, field.DBName)
			methodTableColumnsPtx.Println("\t\t", fmt.Sprintf("%s:%s,", structFieldName, valueColumnName))
		} else {
			valueColumnName := tern.BFF(c.options.columnsMethodRecvName != "" && c.options.columnsCheckFieldType, func() string {
				return fmt.Sprintf(`%s.Cnm(%s.%s, "%s")`, pkgNameGormCnm, c.options.columnsMethodRecvName, field.Name, field.DBName)
			}, func() string {
				return `"` + field.DBName + `"`
			})
			methodPtx.Println("\t\t", fmt.Sprintf("%s:%s,", structFieldName, valueColumnName))
		}
	}
	structPtx.Println("}")
	if c.options.isGenFuncTableColumns {
		methodTableColumnsPtx.Println("\t}")
	} else {
		methodPtx.Println("\t}")
	}
	methodPtx.Println("}")
	methodTableColumnsPtx.Println("}")

	structCode := strings.TrimSpace(structPtx.String())
	methodCode := strings.TrimSpace(methodPtx.String())
	methodTableColumnsCode := strings.TrimSpace(methodTableColumnsPtx.String())

	zaplog.SUG.Debug("---", "\n", methodCode)
	zaplog.SUG.Debug("---", "\n", methodTableColumnsCode)
	zaplog.SUG.Debug("---", "\n", structCode)
	zaplog.SUG.Debug("---", "\n", neatjsons.S(pkgImports))

	return &GenOutput{
		methodCode:             methodCode,
		methodTableColumnsCode: methodTableColumnsCode,
		structCode:             structCode,
		pkgImports:             pkgImports,
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
