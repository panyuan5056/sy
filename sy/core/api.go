package core

import (
	"sy/core/xdb"
	"sy/core/xfile"

	"github.com/gin-gonic/gin"
)

func InitApiv1(apiv1 *gin.RouterGroup) *gin.RouterGroup {
	{
		apiv1.POST("/db/encode", xdb.Encode)
		apiv1.POST("/db/decode", xdb.Decode)
		apiv1.POST("/file/encode", xfile.Encode)
		apiv1.POST("/file/decode", xfile.Decode)

	}
	return apiv1
}
