package str

import "encoding/json"

func AddInSlice(stringSlice []string, str string) []string {
	return append(stringSlice, str)
}

func RemoveInSlice(stringSlice []string, str string) []string {
	for i, strInternal := range stringSlice {
		if strInternal == str {
			return append(stringSlice[:i], stringSlice[i+1:]...)
		}
	}
	return stringSlice
}

func FindInSlice(stringSlice []string, str string) string {
	var result string = ""
	for _, strInternal := range stringSlice {
		if strInternal == str {
			result = strInternal
			break
		}
	}
	return result
}

func UniqueInSlice(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func SliceToJson(stringSlice []string) string {
	if len(stringSlice) == 0 {
		return "[]"
	}

	jsonData, err := json.Marshal(stringSlice)
	if err != nil {
		return "[]"
	}

	return string(jsonData)
}
