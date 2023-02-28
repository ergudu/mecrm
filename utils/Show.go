package utils

import (
	"fmt"
	"github.com/ergudu/mecrm/model"
)

func Show(u model.User) {
	fmt.Printf("%#v \n", u)
}
