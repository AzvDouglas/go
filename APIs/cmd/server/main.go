package main

import (
	"github.com/AzvDouglas/go/APIs/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
