package main

import (
	//_ "./filters"
	_ "youtube-imitation-backstage2/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.Run()

}
