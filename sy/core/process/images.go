package process

import (
	"bufio"
	"bytes"

	"image/gif"
	"image/jpeg"
	"image/png"

	"os"
	"sy/pkg/logging"

	"github.com/auyer/steganography"
)

type Image struct {
	Dump   string
	Ext    string
	Source string
}

func (i *Image) Decode() string {
	inFile, _ := os.Open(i.Dump)
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	img, _ := png.Decode(reader)
	if i.Ext == "jpg" || i.Ext == "jpeg" {
		img, _ = jpeg.Decode(reader)
	} else if i.Ext == "gif" {
		img, _ = gif.Decode(reader)
	}
	sizeOfMessage := steganography.GetMessageSizeFromImage(img)
	msg := steganography.Decode(sizeOfMessage, img)
	return string(msg)
}

func (i *Image) Encode(body string) bool {
	inFile, _ := os.Open(i.Source)
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	img, _ := png.Decode(reader)
	if i.Ext == "jpg" || i.Ext == "jpeg" {
		img, _ = jpeg.Decode(reader)
	} else if i.Ext == "gif" {
		img, _ = gif.Decode(reader)
	}
	w := new(bytes.Buffer)
	err := steganography.Encode(w, img, []byte(body))
	if err != nil {
		logging.Error("Error Encoding file %v", err.Error())
		return false
	}
	outFile, _ := os.Create(i.Dump)
	w.WriteTo(outFile)
	outFile.Close()
	return true
}

func ImageDecode(dump, ext string) string {
	image := &Image{Dump: dump, Ext: ext}
	return image.Decode()
}

func ImageEncode(dump, ext, source, content string) bool {
	image := &Image{Dump: dump, Ext: ext, Source: source}
	return image.Encode(content)
}
