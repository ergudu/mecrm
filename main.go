package main

import (
	"github.com/ergudu/mecrm/model"
	"github.com/ergudu/mecrm/utils"
)

func main() {

	u := model.User{
		Id:   1,
		Name: "HEHEGUNIANG",
		Age:  18,
	}

	utils.Show(u)
}
