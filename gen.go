package gormcngen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/zaplog"
	"gorm.io/gorm/schema"
)

type Config struct {
	sch         *schema.Schema
	nmClassName string
	clsFuncName string
}

func NewConfig(sch *schema.Schema, nmClassName string, clsFuncName string) *Config {
	return &Config{
		sch:         sch,
		clsFuncName: clsFuncName,
		nmClassName: nmClassName,
	}
}

type Options struct {
	IsSubClassExportable bool //根据配置生成非导出的 exampleColumns 或者可导出的 ExampleColumns，通常非导出已经是够用的
}

func NewConfigXObject(dest interface{}, options *Options) *Config {
	sch := done.VCE(schema.Parse(dest, &sync.Map{}, &schema.NamingStrategy{
		SingularTable: false,
		NoLowerCase:   false,
	})).Nice()

	ShowSchemaMessage(sch)

	const classSuffix = "Columns"
	const clsFuncName = "Columns"

	var nmClassName string
	if !options.IsSubClassExportable { //这里不判断options是否非空，默认就是非空（否则调用层写个nil也不太符合预期）
		nmClassName = utils.CvtC0ToLowerString(sch.Name) + classSuffix
	} else {
		nmClassName = sch.Name + classSuffix //这里不用管，通常定义的结构体名称是导出的
	}

	return NewConfig(sch, nmClassName, clsFuncName)
}

type GenResType struct {
	clsFuncCode string
	nmClassCode string
	moreImports map[string]bool
}

func (c *Config) Gen() *GenResType {
	var sch *schema.Schema = c.sch
	var clsFuncName string = c.clsFuncName
	var nmClassName string = c.nmClassName

	pst := utils.NewPTX()

	pst.Println(fmt.Sprintf("type %s struct{", nmClassName))

	cbaType := reflect.TypeOf(gormcnm.ColumnOperationClass{})
	pkgName := filepath.Base(cbaType.PkgPath())

	const align = "   " //让代码对齐的，是3个空格，而不是4个空格，因为打印函数会增加1个空格。由于后面会格式化代码，这里的对齐也只是为了方便观察日志

	pst.Println(align, fmt.Sprintf("%s.%s //继承操作函数，让查询更便捷", pkgName, cbaType.Name()))
	pst.Println(align, "// 模型各个列名和类型:")

	pfu := utils.NewPTX()
	pfu.Println(fmt.Sprintf("func (*%s) %s() *%s {", sch.Name, clsFuncName, nmClassName))
	pfu.Println(fmt.Sprintf("	return &%s{", nmClassName))

	schemaPkgPath := sch.ModelType.PkgPath()

	var moreImports = make(map[string]bool)

	for _, field := range sch.Fields {
		var typName string
		if pkgPath := field.FieldType.PkgPath(); pkgPath == schemaPkgPath { //假如在同一个包里，类型就没必要再用包名
			typName = field.FieldType.Name() //只用类型名即可
		} else {
			if pkgPath != "" {
				moreImports[pkgPath] = true
			}
			typName = field.FieldType.String() //得用完整的名字
		}
		pst.Println(align, field.Name, fmt.Sprintf("%s.ColumnName[%s]", pkgName, typName))

		pfu.Println(align, fmt.Sprintf(`%s:"%s",`, field.Name, field.DBName))
	}

	pfu.Println("	}")
	pfu.Println("}")
	pst.Println("}")

	clsFuncCode := strings.TrimSpace(pfu.String())
	nmClassCode := strings.TrimSpace(pst.String())

	zaplog.LOG.Debug("---")
	fmt.Println(clsFuncCode)
	zaplog.LOG.Debug("---")
	fmt.Println(nmClassCode)
	zaplog.LOG.Debug("---")
	fmt.Println(moreImports)
	zaplog.LOG.Debug("---")

	return &GenResType{
		clsFuncCode: clsFuncCode,
		nmClassCode: nmClassCode,
		moreImports: moreImports,
	}
}

func ShowSchemaMessage(sch *schema.Schema) {
	fmt.Println("---")
	fmt.Println("schema_message", "结构体名称:", sch.Name, "数据表名称:", sch.Table, "模型字段: {") //go结构体名称 和 数据库表名称
	for _, field := range sch.Fields {
		fmt.Println("   ",
			"Go字段名", field.Name, //go结构体成员名称
			"Go类型", field.FieldType, //go的类型
			"DB字段名", field.DBName, //数据表列名
			"DB类型", field.DataType, //数据库的类型
		)
	}
	fmt.Println("}")
	fmt.Println("---")
}
