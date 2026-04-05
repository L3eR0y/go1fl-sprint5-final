package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, info := range dataset {
		err := dp.Parse(info)

		if err != nil {
			log.Println(err)
			continue
		}

		actionInfo, err := dp.ActionInfo()

		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(actionInfo)
		}
	}
}
