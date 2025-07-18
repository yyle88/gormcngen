package gormcngen

import (
	"go/ast"
	"go/token"
	"os"
	"sync/atomic"

	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/rese"
	"github.com/yyle88/sortslice"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_astnode"
	"github.com/yyle88/syntaxgo/syntaxgo_search"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
)

// CodeGenerationConfig defines the configuration of code generation.
// CodeGenerationConfig 是代码生成的配置
type CodeGenerationConfig struct {
	schemas          []*SchemaConfig // Configurations of the model pkg struct schemas. // 自定义模型的 Schema 配置
	methodOutputPath string          // Path where the generated methods will be saved. // 已生成的方法代码保存路径
	structOutputPath string          // Path where the generated struct code be saved. // 已生成的结构体代码保存路径
}

// NewCodeGenerationConfig creates a new instance of CodeGenerationConfig.
// NewCodeGenerationConfig 创建一个新的 CodeGenerationConfig 实例
func NewCodeGenerationConfig(schemas []*SchemaConfig) *CodeGenerationConfig {
	return &CodeGenerationConfig{
		schemas:          schemas,
		methodOutputPath: "",
		structOutputPath: "",
	}
}

// Configs is an alias for CodeGenerationConfig, used for code generation tasks.
// Configs 是 CodeGenerationConfig 的别名，用于代码生成任务
type Configs = CodeGenerationConfig

// NewConfigs initializes a Configs instance based on provided models and options.
// NewConfigs 根据提供的模型和选项初始化 Configs 实例
func NewConfigs(models []interface{}, options *Options, outputPath string) *Configs {
	schemas := make([]*SchemaConfig, 0, len(models))
	for _, object := range models {
		schemas = append(schemas, NewSchemaConfig(object, options)) // Convert models into schema configurations. // 将模型转化为 Schema 配置
	}
	return NewCodeGenerationConfig(schemas).
		WithMethodOutputPath(outputPath).
		WithStructOutputPath(outputPath)
}

// WithMethodOutputPath specifies the output path for method code.
// WithMethodOutputPath 设置方法代码的输出路径
func (cfg *Configs) WithMethodOutputPath(path string) *Configs {
	cfg.methodOutputPath = path
	return cfg
}

// WithStructOutputPath specifies the output path for struct code.
// WithStructOutputPath 设置结构体代码的输出路径
func (cfg *Configs) WithStructOutputPath(path string) *Configs {
	cfg.structOutputPath = path
	return cfg
}

// Generate triggers the code generation process by calling the Gen method.
// Generate 通过调用 Gen 方法触发代码生成过程
func (cfg *Configs) Generate() {
	cfg.Gen()
}

// Gen is the core method responsible for generating code based on the provided schemas.
// Gen 是核心方法，负责根据提供的 schemas 生成代码
func (cfg *Configs) Gen() {
	// Define the elementType struct to store information about code blocks that need editing. // 定义 elementType 结构体，用于存储需要编辑的代码块信息
	type elementType struct {
		sourcePath     string          // Path to the source file containing the code block // 源文件路径
		astNode        ast.Node        // AST node representing the code block // 代码块对应的 AST 节点
		exist          bool            // Flag indicating whether the code block already exists // 标志位，指示代码块是否已存在
		newSourceBlock string          // The new content of the code block to be inserted // 新代码块内容
		pkgImports     map[string]bool // Set of new import statements to be added // 需要添加的导入包集合
	}

	// determine the position-sequence of new code blocks. // 确定新增代码块的顺序位置。
	var newCodeSequence atomic.Int64

	// Create a slice to store all elements that require edits. // 创建一个切片，用于存储所有需要编辑的元素
	var editingElements = make([]*elementType, 0, len(cfg.schemas)*3)

	// Iterate through each schema configuration to generate the corresponding code. // 遍历每个 schema 配置，生成相应的代码
	for _, schemaConfig := range cfg.schemas {
		// Generate code for the current schema. // 生成当前 schema 的代码
		output := schemaConfig.Gen()

		// Handle method-related code logic. // 处理与方法相关的代码逻辑
		if path := cfg.methodOutputPath; path != "" {
			astBundle := rese.P1(syntaxgo_ast.NewAstBundleV4(path))
			astFile, _ := astBundle.GetBundle()

			{
				// Locate the AST definition of the method. // 查找方法的 AST 定义
				methodTypeDeclaration, ok := syntaxgo_search.FindFunctionByReceiverAndName(astFile, schemaConfig.sch.Name, schemaConfig.methodName)
				if ok {
					// If the method already exists, prepare for updating its code block. // 如果方法已存在，准备更新代码块
					editingElements = append(editingElements, &elementType{
						sourcePath:     path,
						astNode:        methodTypeDeclaration,
						exist:          true,
						newSourceBlock: output.methodCode,
						pkgImports:     output.pkgImports,
					})
				} else {
					// 当函数不存在时，是否生成它。前面的当函数已存在时，就必须更新它的内容，避免剩个不维护的函数，被其他逻辑误用
					if schemaConfig.options.isGenNewSimpleColumns {
						// If the method does not exist, prepare for adding a new code block. // 如果方法不存在，准备添加新代码块
						editingElements = append(editingElements, &elementType{
							sourcePath:     path,
							astNode:        syntaxgo_astnode.NewNode(token.Pos(newCodeSequence.Add(1)), 0),
							exist:          false,
							newSourceBlock: output.methodCode,
							pkgImports:     output.pkgImports,
						})
					} else {
						zaplog.LOG.Debug("choose not gen new simple columns function so skip", zap.String("name", schemaConfig.sch.Name+"."+schemaConfig.methodName))
					}
				}
			}

			if schemaConfig.options.isGenFuncTableColumns {
				// Locate the AST definition of the method. // 查找方法的 AST 定义
				methodTypeDeclaration, ok := syntaxgo_search.FindFunctionByReceiverAndName(astFile, schemaConfig.sch.Name, schemaConfig.methodNameTableColumns)
				if ok {
					// If the method already exists, prepare for updating its code block. // 如果方法已存在，准备更新代码块
					editingElements = append(editingElements, &elementType{
						sourcePath:     path,
						astNode:        methodTypeDeclaration,
						exist:          true,
						newSourceBlock: output.methodTableColumnsCode,
						pkgImports:     output.pkgImports,
					})
				} else {
					// If the method does not exist, prepare for adding a new code block. // 如果方法不存在，准备添加新代码块
					editingElements = append(editingElements, &elementType{
						sourcePath:     path,
						astNode:        syntaxgo_astnode.NewNode(token.Pos(newCodeSequence.Add(1)), 0),
						exist:          false,
						newSourceBlock: output.methodTableColumnsCode,
						pkgImports:     output.pkgImports,
					})
				}
			}
		}

		// Handle struct-related code logic. // 处理与结构体相关的代码逻辑
		if path := cfg.structOutputPath; path != "" {
			astBundle := rese.P1(syntaxgo_ast.NewAstBundleV4(path))
			astFile, _ := astBundle.GetBundle()

			// Locate the AST definition of the struct. // 查找结构体的 AST 定义
			// Attempt to find the struct declaration in the AST by its name.
			structTypeDeclaration, ok := syntaxgo_search.FindStructDeclarationByName(astFile, schemaConfig.structName)
			if !ok {
				// If the struct declaration was not found, check if we should ignore the exportable status.
				// 如果没有找到结构体声明，检查是否需要忽略导出状态
				if schemaConfig.options.matchIgnoreExportable {
					// If the struct name is not found, toggle the exportable statue of the struct name.
					// 如果结构体名称未找到，则切换结构体名称的导出性（首字母大小写）
					newStructName := utils.SwitchToggleExportable(schemaConfig.structName)
					// If the toggled struct name is different from the original, try searching again with the new name.
					// 如果切换后的结构体名称与原名称不同，尝试使用新名称再次查找
					if newStructName != schemaConfig.structName {
						// Attempt to find the struct declaration by the new name with the toggled exportable.
						// 尝试使用切换后的名称查找结构体声明
						structTypeDeclaration, ok = syntaxgo_search.FindStructDeclarationByName(astFile, newStructName)
					}
				}
			}
			if ok {
				// If the struct already exists, prepare for updating its code block. // 如果结构体已存在，准备更新代码块
				editingElements = append(editingElements, &elementType{
					sourcePath:     path,
					astNode:        structTypeDeclaration,
					exist:          true,
					newSourceBlock: output.structCode,
					pkgImports:     output.pkgImports,
				})
			} else {
				// If the struct does not exist, prepare for adding a new code block. // 如果结构体不存在，准备添加新代码块
				editingElements = append(editingElements, &elementType{
					sourcePath:     path,
					astNode:        syntaxgo_astnode.NewNode(token.Pos(newCodeSequence.Add(1)), 0),
					exist:          false,
					newSourceBlock: output.structCode,
					pkgImports:     output.pkgImports,
				})
			}
		}
	}

	// Sort the elements based on their existence, with existing code blocks prioritized. // 根据代码块的存在性进行排序，优先处理已存在的代码块
	sortslice.SortVStable[*elementType](editingElements, func(a, b *elementType) bool {
		if a.exist != b.exist {
			return a.exist // Sort existing code blocks to the front // 已存在的代码块排在前面
		} else {
			if a.exist { // If both blocks exist, prioritize the one with the later line number // 如果都已存在，优先排在后面行号更大的代码块
				return a.astNode.Pos() > b.astNode.Pos() // Sort by larger line numbers first // 按行号较大的排在前面
			} else { // If neither block exists, maintain the creation order // 如果都不存在，按创建顺序排序
				return a.astNode.Pos() < b.astNode.Pos() // Keep code blocks created earlier at the front // 保持先创建的代码块排在前面
			}
		}
	})

	// Map file paths to their corresponding source code and imports. // 将文件路径与其对应的源代码和导入包映射
	type sourceAndImportsTuple struct {
		sourceCode []byte          // Source code of the go file // 源文件的完整源代码
		pkgImports map[string]bool // Required import statements // 需要导入的包
	}

	var path2codeMap = map[string]*sourceAndImportsTuple{} // 文件路径与源代码映射

	// Initialize the mapping for each file that requires editing. // 为每个需要编辑的文件初始化映射
	for _, elem := range editingElements {
		if _, ok := path2codeMap[elem.sourcePath]; !ok {
			path2codeMap[elem.sourcePath] = &sourceAndImportsTuple{
				sourceCode: done.VAE(os.ReadFile(elem.sourcePath)).Done(), // Read the existing source code from the file // 读取文件的现有源代码
				pkgImports: map[string]bool{},                             // Initialize with an empty set of imports // 初始化为空的导入包集合
			}
		}
	}

	// Apply changes to the source code for each file based on the collected edit elements. // 根据收集的编辑元素更新每个文件的源代码
	for _, elem := range editingElements {
		tuple := path2codeMap[elem.sourcePath]
		if elem.exist { // If the code block exists, replace it with the new content. // 如果代码块已存在，用新内容替换它
			tuple.sourceCode = syntaxgo_astnode.ChangeNodeCodeSetSomeNewLines(tuple.sourceCode, elem.astNode, []byte(elem.newSourceBlock), 2)
		} else { // If the code block does not exist, append the new content. // 如果代码块不存在，追加新内容
			tuple.sourceCode = append(tuple.sourceCode, byte('\n'), byte('\n'))
			codeBlockBytes := []byte(elem.newSourceBlock)
			tuple.sourceCode = append(tuple.sourceCode, codeBlockBytes...)
		}
		// Add any required imports. // 添加所需的导入包
		for pkgPath := range elem.pkgImports {
			tuple.pkgImports[pkgPath] = true
		}
	}

	// Inject the necessary imports into the source code. // 将必要的导入包注入源代码
	for _, tuple := range path2codeMap {
		options := syntaxgo_ast.NewPackageImportOptions()
		options.SetPkgPaths(maps.Keys(tuple.pkgImports))
		options.SetInferredObject(gormcnm.ColumnOperationClass{})

		tuple.sourceCode = options.InjectImports(tuple.sourceCode)
	}

	// Format the updated source code and write it back to the respective files. // 格式化更新后的源代码，再写回相应的文件
	for absPath, tuple := range path2codeMap {
		newSource := done.VAE(formatgo.FormatBytes(tuple.sourceCode)).Nice()
		done.Done(os.WriteFile(absPath, newSource, 0644))
	}
}
