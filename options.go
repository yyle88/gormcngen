// Package gormcngen: Configuration options for intelligent code generation actions
// Provides fine-grained control of AST-based code generation process
// Supports various generation modes, tag handling, and output customization
//
// gormcngen: 智能代码生成行为的配置选项
// 提供对基于 AST 代码生成过程的精细控制
// 支持各种生成模式、标签处理和输出自定义
package gormcngen

// Options provides fine-grained control of code generation actions
// Configures various aspects of the intelligent code generation process
// Controls struct export status, tag processing, field filtering, and method generation
//
// Options 配置智能代码生成过程的各个方面
// 控制结构体导出可见性、标签处理、字段过滤和方法生成
type Options struct {
	columnClassExportable bool   // Generate exported structures instead of non-exported ones. // 是否生成导出结构或非导出结构
	useTagName            bool   // Use tag names as field names. // 是否使用标签名作为字段名
	tagKeyName            string // Tag that stores field names. // 存储字段名的标签键
	excludeUntaggedFields bool   // Skip fields without tags. // 是否跳过没有标签的字段
	columnsMethodRecvName string // Columns method receiver variable name. // Columns 函数的接受者名称
	columnsCheckFieldType bool   // Columns method checks field types. // Columns 函数是否检查字段类型
	embedColumnOperations bool   // Embed ColumnOperationClass. // 是否嵌套 ColumnOperationClass
	matchIgnoreExportable bool   // Ignore the exportable-case when matching. // 匹配时是否忽略导出性
	isGenFuncTableColumns bool   // Generate the table columns function. // 是否生成表表的列函数
	isGenNewSimpleColumns bool   // Generate the plain columns function. // 是否生成简单的列函数
}

// NewOptions creates a new Options instance with sensible default values
// Initializes configuration optimized for common use cases and best practices
// Returns a pre-configured instance prepared for customization via With methods
//
// NewOptions 创建一个具有合理默认值的 Options 实例
// 初始化针对常见用例和最佳实践优化的配置
// 返回一个预配置的实例，准备通过 With 方法进行自定义
func NewOptions() *Options {
	return &Options{
		columnClassExportable: true,
		useTagName:            false,
		tagKeyName:            "",
		excludeUntaggedFields: false,
		columnsMethodRecvName: "",
		columnsCheckFieldType: false,
		embedColumnOperations: true,
		matchIgnoreExportable: true,
		isGenFuncTableColumns: false,
		isGenNewSimpleColumns: true,
	}
}

func (o *Options) WithColumnClassExportable(columnClassExportable bool) *Options {
	o.columnClassExportable = columnClassExportable
	return o
}

func (o *Options) WithUseTagName(useTagName bool) *Options {
	o.useTagName = useTagName
	return o
}

func (o *Options) WithTagKeyName(tagKeyName string) *Options {
	o.tagKeyName = tagKeyName
	return o
}

func (o *Options) WithExcludeUntaggedFields(excludeUntaggedFields bool) *Options {
	o.excludeUntaggedFields = excludeUntaggedFields
	return o
}

func (o *Options) WithColumnsMethodRecvName(columnsMethodRecvName string) *Options {
	o.columnsMethodRecvName = columnsMethodRecvName
	return o
}

func (o *Options) WithColumnsCheckFieldType(columnsCheckFieldType bool) *Options {
	o.columnsCheckFieldType = columnsCheckFieldType
	return o
}

func (o *Options) WithEmbedColumnOperations(embedColumnOperations bool) *Options {
	o.embedColumnOperations = embedColumnOperations
	return o
}

func (o *Options) WithMatchIgnoreExportable(matchIgnoreExportable bool) *Options {
	o.matchIgnoreExportable = matchIgnoreExportable
	return o
}

func (o *Options) WithIsGenFuncTableColumns(isGenFuncTableColumns bool) *Options {
	o.isGenFuncTableColumns = isGenFuncTableColumns
	return o
}

func (o *Options) WithIsGenNewSimpleColumns(isGenNewSimpleColumns bool) *Options {
	o.isGenNewSimpleColumns = isGenNewSimpleColumns
	return o
}
