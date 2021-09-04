package util

import (
	"github.com/spf13/cast"
)

func Col2row(data map[string][]interface{}) []map[string]interface{} {
	var result []map[string]interface{}
	for name, col := range data {
		if result == nil {
			result = make([]map[string]interface{}, len(col))
		}
		for index, v := range col {
			if result[index] == nil {
				result[index] = map[string]interface{}{}
			}

			result[index][name] = v
		}
	}
	return result
}

func Row2ColRow(data []map[string]interface{}) map[string][]interface{} {
	tmp := map[string][]interface{}{}
	for index, row := range data {
		for name, v := range row {
			if tmp[name] == nil {
				tmp[name] = make([]interface{}, len(data))
			}
			tmp[name][index] = v
		}
	}
	return tmp
}

func Row2col(datas map[string][]map[string]interface{}) map[string]map[string][]interface{} {
	result := map[string]map[string][]interface{}{}
	for table, data := range datas {
		tmp := map[string][]interface{}{}
		for index, row := range data {
			for name, v := range row {
				if tmp[name] == nil {
					tmp[name] = make([]interface{}, len(data))
				}
				tmp[name][index] = v
			}
		}
		result[table] = tmp
	}
	return result
}

<<<<<<< HEAD
func In(l []interface{}, c interface{}) bool {
=======
func In(l []string, c string) bool {
>>>>>>> 1246a84248f5686edcd4dc4d6ac91835a37ceb58
	for _, row := range l {
		if row == c {
			return true
		}
	}
	return false
}
<<<<<<< HEAD
=======

func F2S(data map[string]interface{}) map[string]string {
	tmp := map[string]string{}
	for k, v := range data {
		tmp[k] = cast.ToString(v)
	}
	return tmp
}
>>>>>>> 1246a84248f5686edcd4dc4d6ac91835a37ceb58
