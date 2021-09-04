package xfile

import (
	"sy/models"
	"sy/pkg/e"
	"sy/pkg/logging"
	"sy/pkg/setting"

	"fmt"
	"fx/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Encode(c *gin.Context) {
	var form UploadForm
	code := e.INVALID_PARAMS
	if c.ShouldBind(&form) == nil {
		code = e.SUCCESS
		if ext, ok := form.valid(); ok {
			dst := util.ParseFile("1", ext)
			if err := c.SaveUploadedFile(form.Upload, dst); err != nil {
				logging.Error(err.Error())
			} else {
				if content, ok := form.str(dst); ok {
					if util.In(setting.EXT, ext) {
						if util.In(setting.IMAGEEXT, ext) {
							models.Push("2", content)
						} else {
							models.Push("4", content)
						}
					} else {
						code = e.INVALID_PARAMS
						logging.Error(fmt.Sprintf("%s文件格式符合要求", content))
					}
				}
			}
		} else {
			code = e.ERROR_EXT
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})
	c.Abort()
}

func Decode(c *gin.Context) {
	var form UploadForm
	code := e.INVALID_PARAMS
	if c.ShouldBind(&form) == nil {
		code = e.SUCCESS
		if ext, ok := form.valid(); ok {
			dst := util.ParseFile("1", ext)
			if err := c.SaveUploadedFile(form.Upload, dst); err != nil {
				logging.Error(err.Error())
			} else {
				if content, ok := form.str(dst); ok {
					if util.In(setting.EXT, ext) {
						if util.In(setting.IMAGEEXT, ext) {
							models.Push("3", content)
						} else {
							models.Push("5", content)
						}
					} else {
						code = e.INVALID_PARAMS
						logging.Error(fmt.Sprintf("%s文件格式符合要求", content))
					}
				}
			}
		} else {
			code = e.ERROR_EXT
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})
	c.Abort()
}
