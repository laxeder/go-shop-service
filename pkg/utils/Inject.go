package utils

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func Inject(input interface{}, output interface{}) (err error) {

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

	err = mapstructure.Decode(oMap, &output)

	if err != nil {
		return
	}

	return
}

func InjectMap(input any, output interface{}) (err error) {

	data, err := json.Marshal(input)
	if err != nil {

		return
	}

	err = json.Unmarshal(data, output)
	if err != nil {
		return
	}

	return
}

func InjectBytes(data []byte, output interface{}) (err error) {

	if len(data) == 0 {
		return
	}

	err = json.Unmarshal(data, output)

	if err != nil {
		return
	}

	return
}
