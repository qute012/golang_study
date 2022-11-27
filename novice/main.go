package main

import (
	"novice/hello"
	"novice/var-type"
	"novice/uuid-gen"
)

func main()  {
	hello.Hello("Jin")
	var_type.VarTest()
	uuid := ug.Generate()
	hello.Say(uuid)
}