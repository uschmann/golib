package main

import (
	"fmt"

	"github.com/uschmann/golib/utils"
)

func main() {
	p := utils.Person{
		"Andre",
		"Uschmann",
		34,
	}

	fmt.Println(p.Fullname())
}
