package main

import (
	"github.com/alabastida/vartools"
)

func main() {

	a := "declared and not used"
	b := "another declared and not used"
	c := 123

	vartools.Use(a, b, c)
}
