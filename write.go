package gormcngen

import (
	"fmt"
	"go/ast"
	"go/token"
	"sync"

	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"gitlab.yyle.com/golang/uvcode.git/utils_gen"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file"
	"gitlab.yyle.com/golang/uvyyle.git/utils_golang/utils_golang_ast"
	"gitlab.yyle.com/golang/uvyyle.git/utils_map"
	"gitlab.yyle.com/golang/uvyyle.git/utils_sort/utils_sortslice"
	"gitlab.yyle.com/golang/uvyyle.git/utils_strings"
	"gorm.io/gorm/schema"
)

type GenCfg struct {
	sch           *schema.Schema
	getCsFuncName string
	subStructName string
}

type GenCfgs struct {
	cfgs          []*GenCfg
	CsFuncPath    string
	SubStructPath string
}

func NewGenCfgs(dest interface{}, isExportSubClass bool) *GenCfg {
	sch, err := schema.Parse(dest, &sync.Map{}, &schema.NamingStrategy{
		SingularTable: false,
		NoLowerCase:   false,
	})
	done.Done(err)

	csFuncName := "Columns"
	structName := fmt.Sprintf("%sColumns", sch.Name)
	if !isExportSubClass {
		structName = utils_strings.STYLE.CvtC0Lower(structName)
	}

	return &GenCfg{
		sch:           sch,
		getCsFuncName: csFuncName,
		subStructName: structName,
	}
}

func NewGenCfgsXPath(models []interface{}, path string, isExportSubClass bool) *GenCfgs {
	cfgs := make([]*GenCfg, 0, len(models))
	for _, dest := range models {
		cfgs = append(cfgs, NewGenCfgs(dest, isExportSubClass))
	}
	return &GenCfgs{
		cfgs:          cfgs,
		CsFuncPath:    path,
		SubStructPath: path,
	}
}

func (cfgs *GenCfgs) GenWrite() {
	type elemType struct {
		srcPath string          //代码文件路径
		astNode ast.Node        //代码块所在的起止位置
		exist   bool            //代码块是否找到
		newCode string          //新代码块内容
		impsMap map[string]bool //新增的引用部分
	}

	var elems = make([]*elemType, 0, len(cfgs.cfgs)*2)

	for idx, cfg := range cfgs.cfgs {
		codeDefineFunc, codeStructType, moreImportsMap := GenCode(cfg.sch, cfg.getCsFuncName, cfg.subStructName)
		if path := cfgs.CsFuncPath; path != "" {
			astFile, err := utils_golang_ast.NewAstXFilepath(path)
			done.Done(err)

			astFunc, ok := utils_golang_ast.SeekFuncXRecvNameXFuncName(astFile, cfg.sch.Name, cfg.getCsFuncName, false)
			if ok {
				elems = append(elems, &elemType{
					srcPath: path,
					astNode: astFunc,
					exist:   true,
					newCode: codeDefineFunc,
					impsMap: moreImportsMap,
				})
			} else {
				elems = append(elems, &elemType{
					srcPath: path,
					astNode: utils_golang_ast.NewNode(token.Pos(100*idx)+1, 0), //给个假的，做排序用
					exist:   false,
					newCode: codeDefineFunc,
					impsMap: moreImportsMap,
				})
			}
		}
		if path := cfgs.SubStructPath; path != "" {
			astFile, err := utils_golang_ast.NewAstXFilepath(path)
			done.Done(err)

			structDeclsTypes := utils_golang_ast.SeekStructDeclsTypes(astFile)
			structType, ok := structDeclsTypes[cfg.subStructName]
			if ok {
				elems = append(elems, &elemType{
					srcPath: path,
					astNode: structType,
					exist:   true,
					newCode: codeStructType,
					impsMap: moreImportsMap,
				})
			} else {
				elems = append(elems, &elemType{
					srcPath: path,
					astNode: utils_golang_ast.NewNode(token.Pos(100*idx)+2, 0), //给个假的，做排序用
					exist:   false,
					newCode: codeStructType,
					impsMap: moreImportsMap,
				})
			}
		}
	}
	//其实不同文件中的操作，合在一起排序，有一些取巧的成分，但认为这样做比较简单，因此没有严格区分文件，而是先排序再分文件的
	utils_sortslice.NewSortSlice[*elemType](elems, func(a, b *elemType) bool {
		if a.exist != b.exist {
			return a.exist //认为已存在的要放在前面，而不存在的要放在后面，毕竟都是可以随便补充的
		} else {
			if a.exist { //都存在时，优先替换后面的
				return a.astNode.Pos() > b.astNode.Pos() //让行号更大的放在前面，这样最大的向最小的替换，就不会影响前面的行号
			} else { //都不存在时，按照创建的顺序追加
				return a.astNode.Pos() < b.astNode.Pos() //当不存在时，保持先来的排在前面，序号大的排在后面
			}
		}
	}).Sort()

	var sourcesMap = map[string][]byte{}
	var importsMap = map[string]map[string]bool{}
	for _, elem := range elems {
		if _, ok := sourcesMap[elem.srcPath]; !ok {
			sourcesMap[elem.srcPath] = utils_file.READER.Must().Bytes(elem.srcPath)
			importsMap[elem.srcPath] = map[string]bool{} //这里只需要初始化个空的就行
		}
	}

	for _, elem := range elems {
		source := sourcesMap[elem.srcPath]
		if elem.exist {
			source = utils_golang_ast.ChangeNodeBytesXNewLines(source, elem.astNode, []byte(elem.newCode), 2)
		} else {
			source = append(source, byte('\n'), byte('\n'))
			codeBlockBytes := []byte(elem.newCode)
			source = append(source, codeBlockBytes...)
		}
		sourcesMap[elem.srcPath] = source
		mp := importsMap[elem.srcPath]
		for pkgPath := range elem.impsMap {
			mp[pkgPath] = true //这里只需要追加就行
		}
	}

	for filename, source := range sourcesMap {
		source = utils_golang_ast.AddImportsToSource3(
			filename,
			source,
			utils_map.Keys(importsMap[filename]),
			[]any{gormcnm.ColumnBaseFuncClass{}},
		)
		utils_gen.WriteBytes(filename, source)
	}
}
