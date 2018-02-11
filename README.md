package main

import (
	"github.com/alabastida/vartools"
)

func main() {

	a := "declared and not used"
	b := "another declared and not used"
	c := 123

	// 开发调试的时候所有被提示未使用的变量都放在这里，调试完成后把这行删掉，再编译并删除未使用的变量
	vartools.Use(a, b, c)
}
