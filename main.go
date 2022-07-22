package main

import (
	"go-admin/router"
)

func main() {
	r := router.InitRouter()

	r.Run(":5000")
}