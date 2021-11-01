package main

import (
	"github.com/cuilan/file-index/components"
	"github.com/cuilan/file-index/service"
)

func init() {
	components.InitConfiger()
}

func main() {
	service.Index()
}
