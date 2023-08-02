package gdk_test

import "github.com/researchlab/gdk"

// 可以通过同一个logId 追踪log链路
func ExampleLtx() {
	ltx := gdk.NewLtx().Set("example")
	foo(ltx)
}

func foo(ltx *gdk.Ltx) {
	ltx.Warnf("[warn] num:%d", 1)
	boo(ltx)
}

func boo(ltx *gdk.Ltx) {
	ltx.Warnf("[error] num:%d", 1)
}
