package xfile

import (
	"fx/models"
	"fx/pkg/e"
	"fx/pkg/logging"

	"net/http"

	"fx/pkg/util"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	var form UploadForm
	code := e.INVALID_PARAMS
	data := map[string]int{}
	if c.ShouldBind(&form) == nil {
		code = e.SUCCESS
		if _, ok := form.valid(); ok {
			dst := util.ParseFile("1", form.Upload.Filename)
			if err := c.SaveUploadedFile(form.Upload, dst); err != nil {
				logging.Error(err.Error())
			} else {
				if content, ok := form.str(dst); ok {
					data["id"] = models.Push("2", content)
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": data,
	})
	c.Abort()
}
