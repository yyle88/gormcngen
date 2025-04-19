package gormcngen

// Options Configuration options of the generation logic.
// Options 用于控制生成行为的配置选项。
type Options struct {
	columnClassExportable bool   // Whether to generate exported or non-exported structures. // 是否生成导出结构或非导出结构
	useTagName            bool   // Whether to use tag names as field names. // 是否使用标签名作为字段名
	tagKeyName            string // Tag key to store field names. // 存储字段名的标签键
	excludeUntaggedFields bool   // Whether to skip fields without tags. // 是否跳过没有标签的字段
	columnsMethodRecvName string // Columns method receiver variable name. // Columns 函数的接受者名称
	columnsCheckFieldType bool   // Whether the Columns method checks field types. // Columns 函数是否检查字段类型
	embedColumnOperations bool   // Whether to embed ColumnOperationClass. // 是否嵌套 ColumnOperationClass
	matchIgnoreExportable bool   // Whether to ignore the exportable-case when matching. // 匹配时是否忽略导出性
	isGenFuncTableColumns bool   // Whether to generate the table columns function. // 是否生成表表的列函数
	isGenNewSimpleColumns bool   // Whether to generate the plain columns function. // 是否生成简单的列函数
}

// NewOptions creates a new Options instance with default values.
// NewOptions 用于创建一个具有默认值的 Options 实例。
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
