package shared

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func Inject(input interface{}, output interface{}) interface{} {
	iMap := structs.Map(input)
	oMap := structs.Map(output)

	for k, v := range iMap {

		if fmt.Sprintf("%T", oMap[k]) == "<nil>" {
			continue
		}

		if v == nil || v == 0 || v == "" || v == false || fmt.Sprintf("%v", v) == "" {
			continue
		}

		oMap[k] = v
	}

	err := mapstructure.Decode(oMap, &output)

	if err != nil {
		return output
	}

	return output
}
