# gormcngen 目的是让gorm的硬编码减少些

你需要首先看这个项目 [gormcnm](https://github.com/yyle88/gormcnm)

假设你就是从那个项目过来的，已经深刻理解这个项目的作用，就是生成那个项目 [gormcnm](https://github.com/yyle88/gormcnm) 需要的列定义代码。

接下来告诉你如何使用。

其实内容都在: [这个demo目录里面](/internal/demos/demo1)

首先请看这里: [最简单的demo模型](/internal/demos/demo1/models/example.go) 

接着请看这里: [如何生成代码case](/internal/demos/demo1/models/gormcnm.gen_test.go)

生成的新代码: [得到中间代码code](/internal/demos/demo1/models/gormcnm.gen.go)

这是调用逻辑: [运行业务逻辑main](/internal/demos/demo1/main/main.go)

因此最简单的使用的方法就是，直接拷贝 `gormcnm.gen_test.go` 这个文件到你的项目 model/models 目录里，接着把你想要生成的 model 的类对象写上就行。
当然根据情况你需要略微改改测试代码，以及提前创建 `gormcnm.gen.go` 这个空的go文件。
接着运行测试就行（前提是需要安装那些依赖）。

```
go get github.com/yyle88/gormcngen
```

还有这些样例: [其它样例的目录](internal/examples) 里面演示了多个model的时候如何做，和把代码生成在和model相同的文件里怎么做，但我事后看来这些好像无用功，因为连我自己都不想把代码生成到已经有内容的 model 定义文件里（觉得还是独立出来，随时清空再生成就行）。

具体在使用 `gorm` 时如何配合使用，还是看原来的: [gormcnm](https://github.com/yyle88/gormcnm) 里面有较为详细的说明(其实不详细，毕竟只有1个开发者也没有官网，因此能不能用起来就只能靠悟啦，实在是不好意思啊)。

## 其它
单是把这个开源出来，能让我在以后的公司代码里使用，其收益就已经回本啦。

但我为什么不把 [gormcnm](https://github.com/yyle88/gormcnm) 和 [gormcngen](https://github.com/yyle88/gormcngen) 合为同一个项目呢，我想，这或许是因为最初做的时候就是分【工具包】和【生成包】俩包做的，最后开源出来也保留了这种印记。但也未必两个包必然是同时使用的。就这样吧。
