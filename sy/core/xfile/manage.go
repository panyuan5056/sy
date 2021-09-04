package xfile

import (
	"fmt"
	"fx/core/xfile"
	"fx/pkg/util"
	"sy/core/process"

	"github.com/spf13/cast"
)

func XEncode(config map[string]interface{}) bool {
	tmp := util.F2S(config)
	tables, schemas, datas, _ := xfile.FileFind(tmp)
	datax := map[string][]map[string]interface{}{}
	for _, table := range tables {
		data := util.Row2ColRow(datas[table])
		proces := process.Space{Content: "test", Pk: "operate_id", N: 1, Data: data}
		tmp2 := util.Col2row(proces.Encode())
		datax[table] = tmp2
	}
	xfile.FileWrite(cast.ToString(config["dump"]), tables, schemas, datax)
	return true
}

func XDecode(config map[string]interface{}) {
	tmp := util.F2S(config)
	tables, _, datas, _ := xfile.FileFind(tmp)
	for _, table := range tables {
		data := util.Row2ColRow(datas[table])
		proces := process.Space{Content: "test", Pk: "operate_id", N: 1, Data: data}
		x := proces.Decode()
		fmt.Println("debug:", x)
	}

}
