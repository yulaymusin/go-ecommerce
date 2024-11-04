package main

import (
	"github.com/yulaymusin/go-ecommerce/pkg/api"
)

func main() {
	router := api.SetupRouter()
	router.Run()
}
