package gormcngen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/gormcnm"
	"gorm.io/gorm/schema"
)

func Gen(dest interface{}, isExportSubClass bool) string {
	cfg := NewGenCfgs(dest, isExportSubClass)
	codeDefineFunc, codeStructType, moreImportsMap := GenCode(cfg.sch, cfg.getCsFuncName, cfg.subStructName)
	res := codeDefineFunc + "\n" + codeStructType
	fmt.Println(res)
	fmt.Println(moreImportsMap)
	return res
}

func GenCode(sch *schema.Schema, csFuncName string, subStructName string) (string, string, map[string]bool) {
	ShowSchema(sch)
	pta := utils.NewPTX()

	pta.Println(fmt.Sprintf("type %s struct{", subStructName))

	cbaType := reflect.TypeOf(gormcnm.ColumnBaseFuncClass{})
	pkgName := filepath.Base(cbaType.PkgPath())

	pta.Println(fmt.Sprintf("%s.%s //继承操作函数，让查询更便捷", pkgName, cbaType.Name()))
	pta.Println("//模型各个列名和类型:")

	ptx := utils.NewPTX()
	ptx.Println(fmt.Sprintf("func (*%s) %s() *%s {", sch.Name, csFuncName, subStructName))
	ptx.Println(fmt.Sprintf("	return &%s{", subStructName))

	schemaPkgPath := sch.ModelType.PkgPath()

	var moreImportsMap = make(map[string]bool)

	for _, field := range sch.Fields {
		var typName string
		if pkgPath := field.FieldType.PkgPath(); pkgPath == schemaPkgPath { //假如在同一个包里，类型就没必要再用包名
			typName = field.FieldType.Name() //只用类型名即可
		} else {
			if pkgPath != "" {
				moreImportsMap[pkgPath] = true
			}
			typName = field.FieldType.String() //得用完整的名字
		}
		pta.Println(field.Name, fmt.Sprintf("%s.ColumnName[%s]", pkgName, typName))

		ptx.Println(fmt.Sprintf(`%s:"%s",`, field.Name, field.DBName))
	}

	ptx.Println("	}")
	ptx.Println("}")
	pta.Println("}")
	codeDefineFunc := strings.TrimSpace(ptx.String())
	codeStructType := strings.TrimSpace(pta.String())

	fmt.Println("---")
	fmt.Println(codeDefineFunc)
	fmt.Println("---")
	fmt.Println(codeStructType)
	fmt.Println("---")
	fmt.Println(moreImportsMap)
	fmt.Println("---")

	return codeDefineFunc, codeStructType, moreImportsMap
}

func ShowSchema(sch *schema.Schema) {
	fmt.Println("---")
	fmt.Println("结构体名称", sch.Name)  //go结构体名称
	fmt.Println("数据表名称", sch.Table) //数据库表名称
	for _, field := range sch.Fields {
		fmt.Println(
			"Go字段名", field.Name, //go结构体成员名称
			"Go类型", field.FieldType, //go的类型
			"DB字段名", field.DBName, //数据表列名
			"DB类型", field.DataType, //数据库的类型
		)
	}
	fmt.Println("---")
}
