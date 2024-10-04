# gormcngen 目的是让gorm的硬编码减少些

你需要首先看这个项目 [gormcnm](https://github.com/yyle88/gormcnm)

假设你就是从那个项目过来的，已经深刻理解这个项目的作用，就是生成那个项目 [gormcnm](https://github.com/yyle88/gormcnm) 需要的列定义代码。

接下来告诉你如何使用。

其实内容都在: [这个demo目录里面](/internal/demos/demo1)

首先请看这里: [最简单的demo模型](/internal/demos/demo1/models/example.go) 

接着请看这里: [如何生成代码case](/internal/demos/demo1/models/gormcnm.gen_test.go)

生成的新代码: [得到中间代码code](/internal/demos/demo1/models/gormcnm.gen.go)

这是调用逻辑: [运行业务逻辑main](/internal/demos/demo1/main/main.go)

因此最简单的使用的方法就是，直接拷贝 [模型目录](/internal/demos/demo1/models) 里面的 `gormcnm.gen_test.go` 这个文件到你的项目 model/models 目录里，接着把你想要生成的 model 的类对象写上就行。
当然根据情况你需要略微改改测试代码（比如修改包名 models 改为你的模型包名），接着运行这个测试文件即可得到新代码。
出于安全考虑，在运行测试时为防止写错文件，代码中限制必须找到 `gormcnm.gen_test.go` 测试文件对应的源文件 `gormcnm.gen.go` 以后才能往里面写新代码，因此根据需要可以手动创建这个新代码文件。
接着运行测试就行（前提是需要安装那些依赖）。
```
go get github.com/yyle88/gormcngen
```

## 样例
还有这些样例: [其它样例的目录](internal/examples) 里面演示了多个model的时候如何做，和把代码生成在和model相同的文件里怎么做，但我事后看来这些好像无用功，因为连我自己都不想把代码生成到已经有内容的 model 定义文件里（觉得还是独立出来，随时清空再生成就行）。

## 彩蛋
在这个样例中还有中文编码。

### 使用中文编码
这个样例 [中文样例](internal/examples/example4/example4usage/example4usage_test.go) 使用中文编码，只是个简答的探索。其实我自己早已用起来了中文编码这部分，但是绝大多数人依然认为把英语学好很重要，但是人生苦短，早点把业务做出来然后暴富也是更好的选择。

### 使用母语编码
当然也可以使用其它国家的语言，使用母语非常重要，便于书写和理解，有利于加快开发进度。假如单个开发者使用英文写单个项目的极限代码量是5万行，则使用母语应该能再翻几倍，这样复杂的业务逻辑也就变得很轻松啦，特别适合在业余时间注意力不集中时开发，也很适合断断续续的抽空做。

## 补充
具体在使用 `gorm` 时如何配合使用，还是看原来的: [gormcnm](https://github.com/yyle88/gormcnm) 里面有较为详细的说明(其实不详细，毕竟只有1个开发者也没有官网，因此能不能用起来就只能靠悟啦，实在是不好意思啊)。

## 其它
单是把这个开源出来，能让我在以后的公司代码里使用，其收益就已经回本啦。

但我为什么不把 [gormcnm](https://github.com/yyle88/gormcnm) 和 [gormcngen](https://github.com/yyle88/gormcngen) 合为同一个项目呢，我想，这或许是因为最初做的时候就是分【工具包】和【生成包】俩包做的，最后开源出来也保留了这种印记。但也未必两个包必然是同时使用的。就这样吧。
