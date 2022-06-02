package split

import "strings"

func Split(str string, sep string) []string {
	var res []string
	index := strings.Index(str, sep)
	for index >= 0 {
		res = append(res, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	//index此时为-1，已经out of range [-1:]
	// res = append(res, str[index:])
	res = append(res, str)
	return res
}
