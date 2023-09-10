package main

import (
	"flag"

	"latihan_sqlc/interface/api/router"
)

var loadAppConfig *string

func init() {
	loadAppConfig = flag.String("c", ".env", "insert your configuration file")
	flag.Parse()
}

func main() {
	routers := router.InitializeRouter(loadAppConfig)
	routers.Listen(":9000")
}
