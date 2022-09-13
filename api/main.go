package main

import (
	"admin/controllers"
	"admin/dao"
	"admin/utils"
	"flag"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var cmdOption = struct {
	migrate  *bool
	process  *bool
	customer *bool

	help *bool
}{}

func init() {
	cmdOption.help = flag.Bool("h", false, "show help")

	cmdOption.migrate = flag.Bool("m", false, "migrate database")
}

func main() {
	cmd()

	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, controllers.SessionKey)

	r.Use(gin.Logger(), gin.Recovery(), cors.New(config))

	controllers.RouterRegist(r)
	r.Run(utils.GetConfig().Port)
}

func cmd() {
	flag.Parse()

	if *cmdOption.help {
		flag.Usage()
		os.Exit(0)
	}

	if *cmdOption.migrate {
		dao.Migrate()
		os.Exit(0)
	}
}
