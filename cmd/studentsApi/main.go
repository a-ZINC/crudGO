package main

import (
	"fmt"
	"github.com/a-ZINC/crudGO/internal/config"
)

func main() {
	fmt.Println("Hello, world!")
	cfg := config.LoadConfig()
	fmt.Println(cfg);
}
