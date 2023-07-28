package main

import (
	"fmt"

	"github.com/g-code99/learning-path/week1/contador-palabras/pkg/utils"
)

func main() {
	fmt.Println(utils.FilterDir("./"))
	fmt.Println(utils.ReadAndCountWordAndLines("test.txt"))
}
