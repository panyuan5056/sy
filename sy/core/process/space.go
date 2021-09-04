package process

import (
	"fmt"
	"fx/pkg/logging"
	"fx/pkg/util"
	"sort"
	"strings"

	"github.com/spf13/cast"
)

type Space struct {
	Content string                   //加密内容
	Pk      string                   //加密主键
	N       int                      //密度
	Data    map[string][]interface{} //数据
}

func (self *Space) formatIndex(key string, rows []interface{}) map[int]int {
	result := map[int]int{}
	for index, row := range rows {
		body := cast.ToString(row)
		hash := cast.ToInt64(util.BKDRHash(fmt.Sprintf("%s%s", body, key)))
		if hash%cast.ToInt64(self.N) == 0 {
			result[index] = cast.ToInt(hash % cast.ToInt64(len(key)))
		}
	}
	return result
}

func (self *Space) addBit(index map[int]int, sheet []interface{}, key string) []interface{} {
	for i, k := range index {
		b := cast.ToString(key[k])
		body := strings.Replace(cast.ToString(sheet[k]), " ", "", -1)
		if b == "49" {
			sheet[i] = fmt.Sprintf("%s  ", body)
		} else if b == "48" {
			sheet[i] = fmt.Sprintf("%s ", body)
		}
	}
	return sheet
}

func (self *Space) Encode() map[string][]interface{} {
	key := util.StringToBin(util.Md5(self.Content))
	if pkc, ok := self.Data[self.Pk]; ok {
		index := self.formatIndex(key, pkc)
		for k, sheet := range self.Data {
			if k != self.Pk {
				self.Data[k] = self.addBit(index, sheet, key)
			}
		}
	} else {
		logging.Error(fmt.Sprintf("pk not found:%s", self.Pk))
	}
	return self.Data
}

func (self *Space) mergeStr(data map[int][]string) string {
	tmp := make([]string, len(data))
	keys := []int{}
	for index, _ := range data {
		keys = append(keys, index)
	}
	sort.Ints(keys)
	for index, key := range keys {
		row := data[key]
		a, b := 0, 0
		for _, tag := range row {
			if tag == "1" {
				a = a + 1
			} else if tag == "0" {
				b = b + 1
			}
		}
		if a > b {
			tmp[index] = "1"
		} else if b > a {
			tmp[index] = "0"
		} else {
			tmp[index] = "x"
		}
	}
	return strings.Join(tmp, "")
}

func (self *Space) checkBit(index map[int]int, sheet []interface{}) string {
	var tmp = map[int][]string{}
	for k, v := range index {
		if tmp[v] == nil {
			tmp[v] = make([]string, len(index))
		}
		body := cast.ToString(sheet[k])
		if strings.HasSuffix(body, "  ") {
			tmp[v] = append(tmp[v], "1")
		} else if strings.HasSuffix(body, " ") {
			tmp[v] = append(tmp[v], "0")
		} else {
			tmp[v] = append(tmp[v], "x")
		}
	}
	t := self.mergeStr(tmp)
	self.Similarity(t)
	return t

}

func (self *Space) Similarity(t string) {
	key := util.StringToBin(util.Md5(self.Content))
	x := util.Similarity(key, t)
	fmt.Println(x)
}

func (self *Space) Decode() bool {
	key := util.StringToBin(util.Md5(self.Content))
	if pkc, ok := self.Data[self.Pk]; ok {
		index := self.formatIndex(key, pkc)
		for k, sheet := range self.Data {
			if k != self.Pk {
				content := self.checkBit(index, sheet)
				fmt.Println(content)
			}
		}
	} else {
		logging.Error(fmt.Sprintf("pk not found:%s", self.Pk))
	}
	return false
}
