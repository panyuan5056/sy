package core

import (
	"fx/pkg/util"
	"sy/core/process"
	"sy/core/xfile"
	"sy/pkg/e"
	"sy/pkg/setting"

	"github.com/spf13/cast"
)

func WaterMark(category string, config map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{"status": false, "message": ""}
	if category == "1" {
		//数据库加水印

	} else if category == "6" {
		//数据库发现

	} else if category == "2" {
		//图片加
		dst := cast.ToString(config["dst"])
		if ext := util.ParseExt(dst); len(ext) > 0 {
			if util.In(setting.IMAGEEXT, ext) {
				dump := util.ParseFile("2", ext)
				result["status"] = process.ImageEncode(dump, ext, cast.ToString(config["dst"]), cast.ToString(config["content"]))
				result["message"] = e.SUCCESS

			}
		}
	} else if category == "3" {
		//图片解
		dst := cast.ToString(config["dst"])
		if ext := util.ParseExt(dst); len(ext) > 0 {
			if util.In(setting.IMAGEEXT, ext) {
				result["status"] = true
				result["message"] = process.ImageDecode(dst, ext)
			}
		}
	} else if category == "4" {
		//文件加
		xfile.XEncode(config)

	} else if category == "5" {
		//文件解
		xfile.XDecode(config)

	}
	return result
}
