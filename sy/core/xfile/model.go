package xfile

import (
	"encoding/json"

	"mime/multipart"
	"strings"
	"sy/pkg/logging"
	"sy/pkg/setting"
)

type UploadForm struct {
	Header  string                `form:"header"`
	Upload  *multipart.FileHeader `form:"upload" binding:"required"`
	Name    string                `form:"name"`
	Dump    string                `form:"dump"`
	Content string                `form:"content" json:"dump" `
}

func (u *UploadForm) valid() (string, bool) {
	filenames := strings.Split(u.Upload.Filename, ".")
	if len(filenames) > 1 {
		ext := strings.ToLower(filenames[len(filenames)-1])
		for _, ext2 := range setting.EXT {
			if ext == ext2 {
				return ext, true
			}
		}
	}
	return "", false
}
func (u *UploadForm) str(dst string) (string, bool) {
	c := map[string]string{"name": u.Name, "header": u.Header, "dst": dst, "dump": u.Dump, "content": u.Content}
	b, err := json.Marshal(c)
	if err != nil {
		logging.Error(err.Error())
		return "", false
	}
	return string(b), true
}
