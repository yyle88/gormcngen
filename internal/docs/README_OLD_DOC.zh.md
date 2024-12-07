# gormcngen 给 GORM 模型提供 Columns() 获取列名的函数

在使用本工具之前，请先了解这个项目：[gormcnm](https://github.com/yyle88/gormcnm)。

假设您已理解 `gormcnm` 的作用，接下来我们将介绍如何使用 `gormcngen`。

## 快速上手

使用指南内容可参考：[demo 目录](../../internal/demos/demo1)。

1. **查看示例模型**  
   示例模型代码：[example.go](../../internal/demos/demo1/demo1models/example.go)

2. **了解生成代码的流程**  
   测试生成代码：[gormcnm.gen_test.go](../../internal/demos/demo1/demo1models/gormcnm.gen_test.go)

3. **查看生成结果**  
   生成的中间代码：[gormcnm.gen.go](../../internal/demos/demo1/demo1models/gormcnm.gen.go)

4. **运行业务逻辑**  
   主程序逻辑：[main.go](../../internal/demos/demo1/main.go)

### 最简单的使用方式

直接复制 [models 目录](../../internal/demos/demo1/demo1models) 下的 `gormcnm.gen_test.go` 文件到您的项目 `model/models` 目录。然后，将需要生成的模型类写入测试文件中即可。

**注意事项：**
- 根据项目需求修改测试文件中的包名，例如将 `models` 替换为您的模型包名。
- 运行测试前，需确保目标代码文件（`gormcnm.gen.go`）已存在；若不存在，请手动创建空文件，以避免误写入其他文件。
- 安装依赖后即可运行测试生成代码：
   ```bash
   go get github.com/yyle88/gormcngen
   ```

---

## 样例参考

更多使用示例可见：[examples 目录](../../internal/examples)。其中包括多个模型的生成方式，以及将生成代码与模型代码合并到同一文件的演示。  
不过，建议将生成代码单独存放，便于清理和重新生成，而不是将其与模型定义文件混在一起。

---

## 彩蛋功能

### 使用中文编码

示例：[中文编码样例](../../internal/examples/example4/example4usage/example4usage_test.go) 展示了如何使用中文进行编码。这是一个简单的探索。  
虽然很多人认为学好英文很重要，但对于开发者而言，快速完成业务并实现财务自由或许更加现实。

### 使用母语编码

您也可以使用其他语言编写代码。  
使用母语不仅有助于提升效率，还能让复杂的业务逻辑更轻松实现。  
例如，当单个开发者使用英文写项目的代码极限为 5 万行时，使用母语可能将这一限制提升几倍，非常适合在碎片化时间内进行开发。

---

## 使用补充

在具体结合 `gorm` 使用时，请参考：[gormcnm 项目](https://github.com/yyle88/gormcnm)。虽然项目说明可能不够详细，但可以通过实践逐步掌握。

---

## 其它说明

将此工具开源后，我可以更方便地在公司项目中使用它，收益已远超预期。

至于为什么不将 [gormcnm](https://github.com/yyle88/gormcnm) 和 [gormcngen](https://github.com/yyle88/gormcngen) 合并为一个项目？原因在于最初设计时工具包和生成包是分开的，开源时也延续了这一结构。当然，这两个工具并非必须一起使用，因此保持它们独立使用是最合适的。

---
