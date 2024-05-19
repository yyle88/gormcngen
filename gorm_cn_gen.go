package gormcngen

import (
	"go/ast"
	"go/token"
	"os"

	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/slicesort"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
)

type Configs struct {
	configs          []*Config
	writeClsFuncPath string
	writeNmClassPath string
}

func NewConfigsXPath(models []interface{}, path string, isSubClassExportable bool) *Configs {
	cfgs := make([]*Config, 0, len(models))
	for _, dest := range models {
		cfgs = append(cfgs, NewConfigXObject(dest, isSubClassExportable))
	}
	return &Configs{
		configs:          cfgs,
		writeClsFuncPath: path,
		writeNmClassPath: path,
	}
}

func (cs *Configs) Gen() {
	type elemType struct {
		srcPath     string          //代码文件路径
		astNode     ast.Node        //代码块所在的起止位置
		exist       bool            //代码块是否找到
		newSrcBlock string          //新代码块内容
		moreImports map[string]bool //新增的引用部分
	}

	var elems = make([]*elemType, 0, len(cs.configs)*2) //因为是类型和函数两个操作，这里*2

	for idx, cfg := range cs.configs {
		res := cfg.Gen()
		if path := cs.writeClsFuncPath; path != "" {
			astFile, err := syntaxgo_ast.NewAstXFilepath(path)
			done.Done(err)

			astFunc, ok := syntaxgo_ast.SeekFuncXRecvNameXFuncName(astFile, cfg.sch.Name, cfg.clsFuncName)
			if ok {
				elems = append(elems, &elemType{
					srcPath:     path,
					astNode:     astFunc,
					exist:       true,
					newSrcBlock: res.clsFuncCode,
					moreImports: res.moreImports,
				})
			} else {
				elems = append(elems, &elemType{
					srcPath:     path,
					astNode:     syntaxgo_ast.NewNode(token.Pos(100*idx)+1, 0), //给个假的，做排序用
					exist:       false,
					newSrcBlock: res.clsFuncCode,
					moreImports: res.moreImports,
				})
			}
		}
		if path := cs.writeNmClassPath; path != "" {
			astFile, err := syntaxgo_ast.NewAstXFilepath(path)
			done.Done(err)

			structDeclsTypes := syntaxgo_ast.SeekMapStructNameDeclsTypes(astFile)
			structType, ok := structDeclsTypes[cfg.nmClassName]
			if ok {
				elems = append(elems, &elemType{
					srcPath:     path,
					astNode:     structType,
					exist:       true,
					newSrcBlock: res.nmClassCode,
					moreImports: res.moreImports,
				})
			} else {
				elems = append(elems, &elemType{
					srcPath:     path,
					astNode:     syntaxgo_ast.NewNode(token.Pos(100*idx)+2, 0), //给个假的，做排序用
					exist:       false,
					newSrcBlock: res.nmClassCode,
					moreImports: res.moreImports,
				})
			}
		}
	}
	//其实不同文件中的操作，合在一起排序，有一些取巧的成分，但认为这样做比较简单，因此没有严格区分文件，而是先排序再分文件的
	slicesort.SortVStable[*elemType](elems, func(a, b *elemType) bool {
		if a.exist != b.exist {
			return a.exist //认为已存在的要放在前面，而不存在的要放在后面，毕竟都是可以随便补充的
		} else {
			if a.exist { //都存在时，优先替换后面的
				return a.astNode.Pos() > b.astNode.Pos() //让行号更大的放在前面，这样最大的向最小的替换，就不会影响前面的行号
			} else { //都不存在时，按照创建的顺序追加
				return a.astNode.Pos() < b.astNode.Pos() //当不存在时，保持先来的排在前面，序号大的排在后面
			}
		}
	})

	type srcTuple struct {
		source      []byte          //某个文件的完整源码
		moreImports map[string]bool //需要新增的引用包们
	}

	var srcMap = map[string]*srcTuple{} //路径和内容
	for _, elem := range elems {
		if _, ok := srcMap[elem.srcPath]; !ok {
			srcMap[elem.srcPath] = &srcTuple{
				source:      done.VAE(os.ReadFile(elem.srcPath)).Done(), //读取源代码内容
				moreImports: map[string]bool{},                          //这里只需要初始化个空的就行，稍后再补充内容
			}
		}
	}

	for _, elem := range elems {
		srcNode := srcMap[elem.srcPath]
		if elem.exist { //假如存在就替换它，替换代码块的全部内容
			srcNode.source = syntaxgo_ast.ChangeNodeBytesXNewLines(srcNode.source, elem.astNode, []byte(elem.newSrcBlock), 2)
		} else { //假如不存在就追加它，把内容追加到文件末尾
			srcNode.source = append(srcNode.source, byte('\n'), byte('\n'))
			codeBlockBytes := []byte(elem.newSrcBlock)
			srcNode.source = append(srcNode.source, codeBlockBytes...)
		}
		for pkgPath := range elem.moreImports {
			srcNode.moreImports[pkgPath] = true //这里只需要追加就行
		}
	}

	for absPath, srcNode := range srcMap {
		source := syntaxgo_ast.AddImports(
			srcNode.source,
			&syntaxgo_ast.PackageImportOptions{
				Packages:   utils.GetMapKeys(srcNode.moreImports),
				UsingTypes: nil,
				Objects:    []any{gormcnm.ColumnBaseFuncClass{}},
			},
		)
		newSource := done.VAE(formatgo.FormatBytes(source)).Nice()
		done.Done(utils.WriteFile(absPath, newSource))
	}
}
