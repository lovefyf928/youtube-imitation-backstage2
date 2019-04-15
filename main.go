package main

import (
	"github.com/astaxie/beego"
	//_ "./filters"
	_ "youtube-imitation-backstage2/routers"
)

func main() {

	beego.Run()

}
