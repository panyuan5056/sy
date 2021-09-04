package xfile

import (
	"encoding/csv"
	"fmt"

	"fx/pkg/logging"
	"fx/pkg/util"

	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/extrame/xls"
	"github.com/spf13/cast"
)

type Manage struct {
	Header string
	Name   string
	Ext    string
	Dst    string
}

func (m *Manage) parseData(records [][]string) []map[string]interface{} {
	columns := []string{}
	results := []map[string]interface{}{}
	if len(records) > 0 {
		if m.Header == "1" {
			columns = records[0]
		} else {
			for i := 0; i < len(records[0]); i++ {
				columns = append(columns, cast.ToString(i))
			}
		}
		for index, row := range records {
			record := map[string]interface{}{}
			if index == 0 && m.Header == "1" {
			} else {
				for index2, col := range columns {
					record[col] = row[index2]
				}
				results = append(results, record)
			}
		}
	}
	return results
}

func (m *Manage) Html() string {
	body := util.ReadFile(m.Dst)
	content := string(body)
	return content
}

func (m *Manage) Xlsx() map[string][]map[string]interface{} {
	results := map[string][]map[string]interface{}{}
	f, err := excelize.OpenFile(m.Dst)
	if err != nil {
		logging.Error(err.Error())
		return results
	}
	for index, sheet := range f.GetSheetList() {
		rows, _ := f.GetRows(sheet)
		tmp := m.parseData(rows)
		name := f.GetSheetName(index)
		results[name] = tmp
	}
	return results
}

func (m *Manage) Xls() map[string][]map[string]interface{} {
	results := map[string][]map[string]interface{}{}
	if xlFile, err := xls.Open(m.Dst, "utf-8"); err == nil {
		for i := 0; i <= xlFile.NumSheets(); i++ {
			if sheet1 := xlFile.GetSheet(i); sheet1 != nil {
				data := [][]string{}
				for i2 := 0; i2 <= cast.ToInt(sheet1.MaxRow); i2++ {
					row := sheet1.Row(i2)
					tmp := []string{}
					for i3 := 0; i3 <= row.LastCol(); i3++ {
						tmp = append(tmp, row.Col(i3))
					}
					data = append(data, tmp)
				}
				results[sheet1.Name] = m.parseData(data)
			}
		}
	}
	return results
}

func (m *Manage) Csv() map[string][]map[string]interface{} {
	results := map[string][]map[string]interface{}{}
	fs, err := os.Open(m.Dst)
	if err != nil {
		logging.Error(err.Error())
	}
	defer fs.Close()
	// 初始化csv-reader
	reader := csv.NewReader(fs)

	// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// 允许懒引号（忘记遇到哪个问题才加的这行）
	reader.LazyQuotes = true
	// 返回csv中的所有内容
	records, read_err := reader.ReadAll()
	if read_err != nil {
		fmt.Println("read_err:", read_err)
		return results
	}

	results["sheet0"] = m.parseData(records)
	return results
}

func (m *Manage) Txt() {
	body := util.ReadFile(m.Dst)
	content := string(body)
	fmt.Println(content)
}

func (m *Manage) Run() map[string][]map[string]interface{} {
	filenames := strings.Split(m.Dst, ".")
	ext := filenames[len(filenames)-1]
	switch ext {
	case "xlsx":
		return m.Xlsx()
	case "csv":
		return m.Csv()
	case "xls":
		return m.Xls()
	case "txt":
		m.Txt()
	default:
		m.Html()
	}
	return map[string][]map[string]interface{}{}
}

func FileFind(config map[string]string) ([]string, map[string][]string, map[string][]map[string]interface{}, map[string]string) {
	manage := &Manage{Header: config["header"], Name: config["name"], Dst: config["dst"], Ext: config["ext"]}
	data := manage.Run()
	tables := []string{}
	schemas := map[string][]string{}
	for table, rows := range data {
		tables = append(tables, table)
		schema := []string{}
		for _, items := range rows {
			for name, _ := range items {
				schema = append(schema, name)
			}
			break
		}
		schemas[table] = schema
	}
	return tables, schemas, data, map[string]string{}
}

func FileWrite(name string, tables []string, schemas map[string][]string, data map[string][]map[string]interface{}) {
	path := util.ParseFile("2", name)
	write := Write{Name: path, Tables: tables, Datas: data, Schemas: schemas}
	write.Dump()
}
