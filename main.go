package main

import (
	"github.com/astaxie/beego"
	//_ "./filters"
	_ "./routers"
)

func main() {

	beego.Run()

}
