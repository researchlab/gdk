gdk 结构设计规则

- 统一访问前缀为 gdk.xxxx 

- 工具函数

一律具体命名名如 DeepCopyForMap(T interface{})T , 引用示例如 gdk.DeepCopyForMap() 

- 结构体变量

首先通过gdk.NewXX() 得到结构体对象， 然后通过结构体对象再调用其业务逻辑如
```golang
type Heap struct{} 

type Heap interface{
    Push(xx){}
}
// 通过返回interface  通过这个interface提供一个函数列表方便查看整体对外开放的接口情况
// 因为有些内部工具函数是不对外开放的， 夹杂在中间 不容易一下看出当前对外开放的接口列表
// 但是 interface{} 变量命名比较麻烦, 理解中的接口应该是用来设计共同函数，松耦合场景的地方用
func NewHeap() HeapImp {return &heap{}}

func (h *heap)Push(xx){}

heap := gdk.NewHeap()  // 然后再由 heap.Push()
```

> 注意: 因为golang 通过首字母来决定结构体变量 是否对外开放， 一般情况下是建议使用小写的， 但是如果后继要在外部解码该结构，或者注入新的信息则不能操作， 所以建议结构体命名一律大写，对外开放权限;

- 第三方包封装如mysql, mq, http.client, response 等

首先通过gdk.NewMysql()得到 mysqlObj 然后 mysqlObj.NewConn() 等操作在继续， 也可以链式 gdk.NewMysql().NewConn()


关于测试

xx_test.go 里面包含 unit_test, benmark test 
xx_example_test.go 包含 示例测试

关于文档 

- 便于生成golang标准文档, 必要时可以通过godoc 架设本地文档服务，快速搜索查看对应的功能和用法

- 便于生成md文档




```shell
#使用go mod
go mod tidy
go mod vendor

#单元测试
go test -race

#压测
time go test -bench=. -run=none
time go test -v -bench=. -cpu=4 -benchtime="10s" -timeout="15s" -benchmem

#代码覆盖率
go test -cover #概览

go test -coverprofile=coverage.out #生成统计信息
go test -v -covermode=count -coverprofile=coverage.out
go tool cover -func=coverage.out #查看统计信息
go tool cover -html=coverage.out #将统计信息转换为html

#性能分析
time go test -timeout 30m -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
go tool pprof profile.out
go tool pprof -http=:8081  profile.out
```
